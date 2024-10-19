package main

import (
	"flag"
	"status/internal/handler"
	"status/internal/status"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// init gin
	r := gin.Default()
	godotenv.Load()
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
	r.Static("/status/resources/", "./front/resources")
	// disable cache for static resources
	r.Use(func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache")
		// if strings.HasPrefix(c.Request.URL.Path, "/resources/") {
		// 	c.Header("Cache-Control", "private, max-age=1")
		// }
		c.Next()
	})
	handler := handler.NewHandler()
	// define routes
	r.GET("/status", handler.Index(config))

	// lezzgo
	r.Run("0.0.0.0:8086")
}
