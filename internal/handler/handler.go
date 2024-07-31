package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"status/internal/service"
	"status/internal/status"

	"github.com/gin-gonic/gin"
)

func Index(config []status.Config) func(*gin.Context) {
	return func(c *gin.Context) {
		healthchecks := make([]service.ServiceHealth, 0)
		for _, serviceConf := range config {
			health := service.ServiceHealth{
				Name: serviceConf.Name,
			}
			if serviceConf.Daemon {
				if err := health.ProvideSystemctlStatus(); err != nil {
					health.Message = err.Error()
				}
			}
			res, err := http.Get("http://" + serviceConf.HealthcheckUrl)
			if err != nil {
				healthchecks = append(healthchecks, service.ServiceHealth{
					Name:    serviceConf.Name,
					Health:  "KO",
					Message: health.Message + "\n" + err.Error(),
				})
				continue
			}
			body, err := io.ReadAll(res.Body)
			if err != nil {
				healthchecks = append(healthchecks, service.ServiceHealth{
					Name:    serviceConf.Name,
					Health:  "KO",
					Message: health.Message + "\n" + err.Error(),
					// Message: err.Error(),
				})
				continue
			}
			err = json.Unmarshal(body, &health)
			if err != nil {
				healthchecks = append(healthchecks, service.ServiceHealth{
					Name:    serviceConf.Name,
					Health:  "KO",
					Message: err.Error(),
				})
				continue
			}
			healthchecks = append(healthchecks, health)
		}
		c.HTML(http.StatusOK, "index.tmpl", healthchecks)
	}
}
