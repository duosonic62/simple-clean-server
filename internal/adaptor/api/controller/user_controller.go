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
	useCase inputboundary.UserUseCase
}

// コンストラクタ
func NewUserController(useCase inputboundary.UserUseCase) UserController {
	return UserControllerImpl{
		useCase: useCase,
	}
}

// ユーザ作成
func (controller UserControllerImpl) CreateUser(ctx domain.Context) {
	data := input.UserAddInputData{}
	ctx.Bind(&data)
	controller.useCase.AddUser(data, ctx)
}
