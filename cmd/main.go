package main

import (
	"github.com/duosonic/simple-clean-server/cmd/injector"
	"github.com/duosonic/simple-clean-server/internal/adaptor/infrastructure/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := router.Router
	var webApp = injector.Initialize()

	// ユーザ登録API
	r.POST("/user", func(context *gin.Context) { webApp.UserController.CreateUser(context) })

	r.Run(":8080")
}
