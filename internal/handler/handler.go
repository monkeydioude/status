package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"status/internal/status"

	"github.com/gin-gonic/gin"
)

func Index(config []status.Config) func(*gin.Context) {
	healthchecks := make([]status.ServiceHealth, 0)
	for _, service := range config {
		res, err := http.Get("http://" + service.HealthcheckUrl)
		if err != nil {
			healthchecks = append(healthchecks, status.ServiceHealth{
				Name:    service.Name,
				Health:  "KO",
				Message: err.Error(),
			})
			continue
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			healthchecks = append(healthchecks, status.ServiceHealth{
				Name:    service.Name,
				Health:  "KO",
				Message: err.Error(),
			})
			continue
		}
		health := status.ServiceHealth{}
		err = json.Unmarshal(body, &health)
		if err != nil {
			healthchecks = append(healthchecks, status.ServiceHealth{
				Name:    service.Name,
				Health:  "KO",
				Message: err.Error(),
			})
			continue
		}
		health.Name = service.Name
		healthchecks = append(healthchecks, health)
	}
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", healthchecks)
	}
}
