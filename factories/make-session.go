package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/user"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/user"
)

func MakeSession(jwtSecret string) controller.CreateSessionController {
	UserRepository := repository.NewUserRepository(db.RunDB())
	CreateSessionUseCase := usecase.NewSessionUseCase(UserRepository)
	CreateSessionController := controller.NewCreateSessionController(CreateSessionUseCase, jwtSecret)

	return CreateSessionController
}
