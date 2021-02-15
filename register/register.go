package register

import (
	"fmt"
	"net/http"

	"github.com/arturoeanton/Yuenyeung/commons"
	"github.com/gin-gonic/gin"
)

// Handler is the interface for add behaviour
type Handler interface {
	Endpoint(c *gin.Context)
	GetPath() string
	GetGroup() string
	GetMethods() []string
	GetRoles() []string
}

var (
	//Catalog is list of the handlers reg
	Catalog = make(map[string]Handler)
	router  *gin.Engine
)

// Register is register in Catalog of handlers
func Register(genericHandler interface{}) {
	h := genericHandler.(Handler)
	Catalog[h.GetPath()] = h
}

func interceptorEndpoint(f func(c *gin.Context), h Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		roles := h.GetRoles()
		if roles != nil {
			// TODO: put validation roles
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "Unauthorized",
				"code":   401,
			})
			return
		}
		f(c)
	}
}

// UpdateRouterWithEndpoints is UpdateRouterWithEndpoints :)
func UpdateRouterWithEndpoints(router *gin.Engine) {
	for k, v := range Catalog {
		group := router.Group(v.GetGroup())
		if v.GetMethods() == nil {
			fmt.Println(v.GetMethods())
			group.Any(k, interceptorEndpoint(v.Endpoint, v))
			continue
		}
		if commons.Contains(v.GetMethods(), http.MethodGet) {
			group.GET(k, interceptorEndpoint(v.Endpoint, v))
		}
		if commons.Contains(v.GetMethods(), http.MethodPost) {
			group.POST(k, interceptorEndpoint(v.Endpoint, v))
		}
		if commons.Contains(v.GetMethods(), http.MethodPut) {
			group.PUT(k, interceptorEndpoint(v.Endpoint, v))
		}
		if commons.Contains(v.GetMethods(), http.MethodDelete) {
			group.DELETE(k, interceptorEndpoint(v.Endpoint, v))
		}
		if commons.Contains(v.GetMethods(), http.MethodPatch) {
			group.PATCH(k, interceptorEndpoint(v.Endpoint, v))
		}
		if commons.Contains(v.GetMethods(), http.MethodHead) {
			group.HEAD(k, interceptorEndpoint(v.Endpoint, v))
		}
		if commons.Contains(v.GetMethods(), http.MethodOptions) {
			group.OPTIONS(k, interceptorEndpoint(v.Endpoint, v))
		}
	}
}

// HandlerBase is the base all Handlers
type HandlerBase struct {
	Path    string
	Group   string
	Methods []string
	Roles   []string
}

// GetPath is implement of interface Handler
func (h *HandlerBase) GetPath() string {
	return h.Path
}

// GetMethods is implement of interface Handler
func (h *HandlerBase) GetMethods() []string {
	return h.Methods
}

// GetGroup is implement of interface Handler
func (h *HandlerBase) GetGroup() string {
	return h.Group
}

// GetRoles is implement of interface Handler
func (h *HandlerBase) GetRoles() []string {
	return h.Roles
}
