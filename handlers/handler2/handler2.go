package handler2

import (
	"net/http"

	"github.com/arturoeanton/Yuenyeung/register"
	"github.com/gin-gonic/gin"
)

// Handler2 is example2
type Handler2 struct {
	register.HandlerBase
}

func init() {
	h := new(Handler2)
	h.Path = "/handler2/:param"
	h.Group = "/v1"
	h.Methods = []string{http.MethodGet}
	register.Register(h)
}

// Endpoint is implement of interface Handler
func (h *Handler2) Endpoint(c *gin.Context) {
	param := c.Param("param")
	c.JSON(200, gin.H{
		"message": param,
		"handler": "2",
	})
}
