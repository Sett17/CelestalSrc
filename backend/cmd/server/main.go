package main

import (
	"celestralsrc/backend/internal/api"
	"celestralsrc/backend/internal/logger"

	"github.com/charmbracelet/log"
	// "github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//go:generate mkdir -p static
//go:generate bash -c "cp ../../../frontend/static/* static/"
//go:generate cp ../../../frontend/dist/main.wasm static/

func main() {
	log.Info("✨ Starting CelestralSrc ✨")
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: logger.Formatter,
		Output:    gin.DefaultWriter,
	}))
	r.Use(gin.Recovery())
    // r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	r.Static("/static", "./static")

    api.SetRoutes(r)

	r.Run(":8080")
}
