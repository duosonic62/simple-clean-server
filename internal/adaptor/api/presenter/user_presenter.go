package presenter

import (
	"github.com/duosonic/simple-clean-server/internal/domain"
	"github.com/duosonic/simple-clean-server/pkg/usecase/output"
)

type UserPresenterImpl struct {

}

func (presenter UserPresenterImpl) OutputAddUser(data output.UserAddOutputData, ctx domain.Context) {
	ctx.JSON(200, data)
}