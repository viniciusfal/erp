package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/user"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/user"
)

func MakeUser() controller.CreateUserController {
	UserRepository := repository.NewUserRepository(db.RunDB())
	CreateUserUseCase := usecase.NewUserUseCase(UserRepository)
	CreateUserController := controller.NewCreateUserController(CreateUserUseCase)

	return CreateUserController
}
