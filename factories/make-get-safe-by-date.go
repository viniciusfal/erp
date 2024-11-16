package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/safe"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/safe"
)

func MakeGetSafesByDate() controller.GetSafesByDateController {
	SafeRepository := repository.NewSafeRepository(db.RunDB())
	GetSafesByDateUseCase := usecase.NewGetSafesByDateUseCase(SafeRepository)
	GetSafesByDateController := controller.NewGetSafesByDateController(GetSafesByDateUseCase)

	return GetSafesByDateController
}
