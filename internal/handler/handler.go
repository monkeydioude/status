package handler

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"status/internal/service"
	"status/internal/status"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	OK   = "OK"
	KO   = "KO"
	NANI = "??"
)

type Handler struct {
	basicAuth service.BasicAuth
}

func NewHandler() *Handler {
	return &Handler{
		basicAuth: service.NewBasicAuth(),
	}
}

func (h *Handler) CheckBasicAuth(r *http.Request, header func(string, string)) error {
	if !h.basicAuth.IsSet {
		return nil
	}
	header("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	l, p, ok := r.BasicAuth()
	if !ok {
		return errors.New("no basic auth")
	}
	usernameHash := sha256.Sum256([]byte(l))
	passwordHash := sha256.Sum256([]byte(p))
	fmt.Printf("l: %s, p: %s, usernameHash: %x, passwordHash: %x, h.basicAuth.Login: %s, h.basicAuth.Password: %s\n", l, p, usernameHash, passwordHash, h.basicAuth.Login, h.basicAuth.Password)
	usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], h.basicAuth.Login[:]) == 1)
	passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], h.basicAuth.Password[:]) == 1)
	if !usernameMatch || !passwordMatch {
		return errors.New("username or password don't match")
	}
	return nil
}

func (h *Handler) Index(config []status.Config) func(*gin.Context) {
	return func(c *gin.Context) {
		if err := h.CheckBasicAuth(c.Request, c.Header); err != nil {
			c.String(401, "Unauthorized: "+err.Error())
			return
		}

		healthchecks := make([]service.ServiceHealth, 0)
		var wg sync.WaitGroup
		// looping through the service list
		for _, serviceConf := range config {
			wg.Add(1)
			go func(serviceConf status.Config) {
				defer wg.Done()
				health := service.ServiceHealth{
					Name:   serviceConf.Name,
					Health: KO,
				}
				// trying to fetch the task's status.
				// works only for linux atm.
				if serviceConf.Daemon {
					if err := health.ProvideSystemctlStatus(); err != nil {
						health.Message = err.Error()
					}
				}
				// prepend with http:// if no scheme is provided
				url := serviceConf.HealthcheckUrl
				if !strings.HasPrefix(url, "http") {
					url = "http://" + serviceConf.HealthcheckUrl
				}
				// requesting the healthcheck endpoint
				res, err := http.Get(url)
				// should actually answering KO if there's an issue with the request
				if err != nil {
					healthchecks = append(healthchecks, service.ServiceHealth{
						Name:    serviceConf.Name,
						Health:  KO,
						Message: health.Message + "\n" + err.Error(),
					})
					return
				}
				// reading the body
				body, err := io.ReadAll(res.Body)
				// should not answer KO if there's an with reading the body.
				if err != nil {
					healthchecks = append(healthchecks, service.ServiceHealth{
						Name:    serviceConf.Name,
						Health:  NANI,
						Message: health.Message + "\n" + err.Error(),
					})
					return
				}
				err = json.Unmarshal(body, &health)
				// should not answer KO if there's an issue with unmarshalling.
				if err != nil {
					healthchecks = append(healthchecks, service.ServiceHealth{
						Name:    serviceConf.Name,
						Health:  NANI,
						Message: health.Message + "\n" + err.Error(),
					})
					return
				}

				// a status code 200 is all that matters to be OK
				if res.StatusCode != 200 {
					health.Health = KO
				} else {
					health.Health = OK
				}
				healthchecks = append(healthchecks, health)
			}(serviceConf)
		}
		wg.Wait()
		// render the HTML template
		c.HTML(http.StatusOK, "index.tmpl", healthchecks)
	}
}
