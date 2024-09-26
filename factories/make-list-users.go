package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func MakeListUsers() controller.ListUserController {
	UserRepository := repository.NewUserRepository(db.RunDB())
	ListUserUseCase := usecase.NewListUserUseCase(UserRepository)
	ListUserController := controller.NewListUserController(ListUserUseCase)

	return ListUserController
}
