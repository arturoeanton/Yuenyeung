package handler_example_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arturoeanton/Yuenyeung/handlers/handler_example"
	"github.com/arturoeanton/Yuenyeung/register"
	"github.com/gin-gonic/gin"
)

func TestHandlerExampleEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(r)

	h := new(handler_example.HandlerExample)
	h.Path = "/handlerExample/:param"
	h.Methods = []string{http.MethodGet, http.MethodPost}
	register.Register(h)
	context.Params = append(context.Params, gin.Param{Key: "param", Value: "test"})
	h.Endpoint(context)
	if r.Result().StatusCode != 200 {
		t.Errorf("Error Status")
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Result().Body)
	jsonResult := buf.String()
	if "{\"handler\":\"Example\",\"message\":\"test\"}" != jsonResult {
		t.Error(jsonResult)
	}

}
