package outputboundary

import (
	"github.com/duosonic/simple-clean-server/internal/domain"
	"github.com/duosonic/simple-clean-server/pkg/usecase/output"
)

type UserPresenter interface {
	OutputAddUser(data output.UserAddOutputData, ctx domain.Context)
}