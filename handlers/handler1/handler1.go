package handler1

import (
	"net/http"

	"github.com/arturoeanton/Yuenyeung/register"
	"github.com/gin-gonic/gin"
)

// Handler1 is example1
type Handler1 struct {
	register.HandlerBase
}

func init() {
	h := new(Handler1)
	h.Path = "/handler1/:param"
	h.Methods = []string{http.MethodGet, http.MethodPost}
	register.Register(h)
}

// Endpoint is implement of interface Handler
func (h *Handler1) Endpoint(c *gin.Context) {
	param := c.Param("param")
	c.JSON(200, gin.H{
		"message": param,
		"handler": "1",
	})
}
