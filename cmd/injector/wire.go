package injector

import (
	"github.com/duosonic/simple-clean-server/internal/adaptor/api/controller"
	"github.com/duosonic/simple-clean-server/internal/adaptor/api/presenter"
	"github.com/duosonic/simple-clean-server/internal/adaptor/infrastructure/factoryimple"
	"github.com/duosonic/simple-clean-server/internal/adaptor/infrastructure/repositoryimple"
	"github.com/duosonic/simple-clean-server/internal/domain/factory"
	"github.com/duosonic/simple-clean-server/internal/usecase/interactor"
)

type WebApp struct {
	UserController controller.UserController
}

func Initialize() *WebApp {
	var idFactory = factoryimple.NewIdFactory()
	var userFactory = factory.NewUserFactory(idFactory)
	var userRepository = repositoryimple.NewUserRpository()
	var userPresenter = presenter.NewUserPresenter()
	var userUseCase = interactor.NewUserUsecase(
		userPresenter,
		userRepository,
		idFactory,
		userFactory,
	)
	var userController = controller.NewUserController(userUseCase)

	var webApp = WebApp{
		UserController: userController,
	}

	return &webApp
}
