package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/user"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/user"
)

func MakeListUsers() controller.ListUserController {
	UserRepository := repository.NewUserRepository(db.RunDB())
	ListUserUseCase := usecase.NewListUserUseCase(UserRepository)
	ListUserController := controller.NewListUserController(ListUserUseCase)

	return ListUserController
}
