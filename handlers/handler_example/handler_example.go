package handler_example

import (
	"net/http"

	"github.com/arturoeanton/Yuenyeung/register"
	"github.com/gin-gonic/gin"
)

// HandlerExample is exampleExample
type HandlerExample struct {
	register.HandlerBase
}

func init() {
	h := new(HandlerExample)
	h.Path = "/handlerExample/:param"
	h.Methods = []string{http.MethodGet, http.MethodPost}
	register.Register(h)
}

// Endpoint is implement of interface Handler
func (h *HandlerExample) Endpoint(c *gin.Context) {
	param := c.Param("param")
	c.JSON(200, gin.H{
		"message": param,
		"handler": "Example",
	})
}
