package register_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arturoeanton/Yuenyeung/register"
	"github.com/gin-gonic/gin"
)

type HandlerTest struct {
	register.HandlerBase
}

var h = new(HandlerTest)

func init() {
	h = new(HandlerTest)
	h.Path = "/test"
	h.Group = "/test"
	h.Roles = []string{"USER"}
	h.Methods = []string{http.MethodGet}
}

func (h *HandlerTest) Endpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
	})
}

func TestHandlerBase(t *testing.T) {
	if h.Path != h.GetPath() {
		t.Errorf("Test Handler GetPath error")
	}
	if h.Group != h.GetGroup() {
		t.Errorf("Test Handler Group error")
	}
	if len(h.GetRoles()) != 1 {
		t.Errorf("Test Handler Roles error (len)")
	}
	if len(h.GetRoles()) == 1 && h.Roles[0] != h.GetRoles()[0] {
		t.Errorf("Test Handler Roles error")
	}

	if len(h.GetMethods()) != 1 {
		t.Errorf("Test Handler Methods error (len)")
	}
	if len(h.GetMethods()) == 1 && h.Methods[0] != h.GetMethods()[0] {
		t.Errorf("Test Handler Methods error")
	}
}

func TestRegister(t *testing.T) {
	register.Register(h)
	if _, ok := register.Catalog[h.Path]; !ok {
		t.Errorf("Test Register error")
	}
}

func TestUpdateRouterWithEndpoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := httptest.NewRecorder()
	_, e := gin.CreateTestContext(r)
	register.UpdateRouterWithEndpoints(e)
}
