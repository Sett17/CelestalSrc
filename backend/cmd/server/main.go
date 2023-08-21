package main

import (
	"celestralsrc/backend/internal/logger"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/gzip"
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
    r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	r.Static("/static", "./static")

	r.POST("/dikka", func(c *gin.Context) {
		c.String(200, "dikka cool")
	})

	r.Run(":8080")
}
