package main

import (
	"flag"
	"status/internal/handler"
	"status/internal/status"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// init gin
	r := gin.Default()

	// flag parsing
	var configPath string
	flag.StringVar(&configPath, "c", "", "/path/to/config.json")
	flag.Parse()

	// load templates
	r.LoadHTMLFiles("front/templates/index.go.tpl")

	// parse the config file
	config, err := status.Parse(configPath)
	if err != nil {
		panic(err)
	}

	// allow to reach resources
	r.Static("/resources/", "./front/resources")
	// disable cache for static resources
	r.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/resources/") {
			c.Header("Cache-Control", "private, max-age=0")
		}
		c.Next()
	})
	// define routes
	r.GET("/status", handler.Index(config))

	// lezzgo
	r.Run("0.0.0.0:8082")
}
