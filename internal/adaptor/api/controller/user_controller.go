package controller

import (
	"github.com/duosonic/simple-clean-server/internal/domain"
	"github.com/duosonic/simple-clean-server/internal/usecase/inputboundary"
	"github.com/duosonic/simple-clean-server/pkg/usecase/input"
)

type SampleResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserController interface {
	CreateUser(c domain.Context)
}

type UserControllerImpl struct {
	service inputboundary.UserUseCase
}

// コンストラクタ
func NewUserController(service inputboundary.UserUseCase) UserController {
	return UserControllerImpl{
		service: service,
	}
}

// ユーザ作成
func (con UserControllerImpl) CreateUser(ctx domain.Context) {
	data := input.UserAddInputData{}
	ctx.Bind(&data)
	con.service.AddUser(data, ctx)
}
