package api

import (
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.Any("/dikka", func(c *gin.Context) {
            c.String(200, "dikka")
        })
        api.GET("/wasm", func(c *gin.Context) {
            c.String(200, `
<script src="static/wasm_exec.js"></script>
<script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("static/main.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
    });
</script>`)
        })
    }
}
