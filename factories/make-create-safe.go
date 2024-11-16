package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/safe"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/safe"
)

func MakeSafe() controller.CreateSafeController {
	SafeRepository := repository.NewSafeRepository(db.RunDB())
	CreateSafeUseCase := usecase.NewCreateSafeUseCase(SafeRepository)
	CreateSafeController := controller.NewCreateSafeController(CreateSafeUseCase)

	return CreateSafeController
}
