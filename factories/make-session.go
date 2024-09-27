package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func MakeSession() controller.CreateSessionController {
	UserRepository := repository.NewUserRepository(db.RunDB())
	CreateSessionUseCase := usecase.NewSessionUseCase(UserRepository)
	CreateSessionController := controller.NewCreateSessionController(CreateSessionUseCase)

	return CreateSessionController
}
