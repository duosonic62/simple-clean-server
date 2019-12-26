package main

import (
	"github.com/duosonic/simple-clean-server/internal/adaptor/api/controller"
	"github.com/duosonic/simple-clean-server/internal/adaptor/infrastructure/router"
)

func main()  {
	r := router.Router
	r.GET("/sample", controller.Sample)

	r.Run(":8080")
}
