package inputboundary

import (
	"github.com/duosonic/simple-clean-server/internal/domain"
	"github.com/duosonic/simple-clean-server/pkg/usecase/input"
)

type UserUseCase interface {
	AddUser(data input.UserAddInputData, ctx domain.Context)
}