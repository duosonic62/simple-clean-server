package controller

import "github.com/gin-gonic/gin"

type SampleResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Sample(ctx *gin.Context) {
	res := SampleResponse{
		Id:   1,
		Name: "hoge",
	}

	ctx.JSON(200, res)
}
