package handler

import (
	"encoding/json"
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

func Index(config []status.Config) func(*gin.Context) {
	return func(c *gin.Context) {
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
