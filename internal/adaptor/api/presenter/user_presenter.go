package presenter

import (
	"github.com/duosonic/simple-clean-server/internal/domain"
	"github.com/duosonic/simple-clean-server/internal/usecase/outputboundary"
	"github.com/duosonic/simple-clean-server/pkg/usecase/output"
)

type UserPresenterImpl struct{}

// コンストラクタ
func NewUserPresenter() outputboundary.UserPresenter {
	return UserPresenterImpl{}
}

func (presenter UserPresenterImpl) OutputAddUser(data output.UserAddOutputData, ctx domain.Context) {
	ctx.JSON(200, data)
}
