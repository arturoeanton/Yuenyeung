package main

import (
	"fmt"

	_ "github.com/arturoeanton/Yuenyeung/handlers/handler1"
	_ "github.com/arturoeanton/Yuenyeung/handlers/handler2"
	"github.com/arturoeanton/Yuenyeung/register"
	"github.com/gin-gonic/gin"
)

func main() {
	addr := "0.0.0.0:8000"
	//gin.SetMode(gin.ReleaseMode)
	fmt.Println("Server:	", addr)
	router := gin.Default()
	register.UpdateRouterWithEndpoints(router)
	router.Run(addr)
}

func init() {
	fmt.Println("Yuenyeung		v0.0.1")
}

// TODO: add rabbitmq integration
// TODO: add redis integration
// TODO: add auth0 integration
// TODO: add mongo integration
// TODO: add gorm integration
