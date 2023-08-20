package main

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

//go:generate mkdir -p static
//go:generate bash -c "cp ../../../frontend/static/* static/"
//go:generate cp ../../../frontend/dist/main.wasm static/

func main() {
    log.Info("✨ Starting CelestralSrc ✨")
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
        c.File("./static/index.html")
	})

    r.Static("/static", "./static")

	r.Run(":8080")
}
