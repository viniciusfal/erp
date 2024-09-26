package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func MakeUser() controller.CreateUserController {
	UserRepository := repository.NewUserRepository(db.RunDB())
	CreateUserUseCase := usecase.NewUserUseCase(UserRepository)
	CreateUserController := controller.NewCreateUserController(CreateUserUseCase)

	return CreateUserController
}
