package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/safe"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/safe"
)

func MakeListSafe() controller.ListSafeController {
	SafeRepository := repository.NewSafeRepository(db.RunDB())
	ListSafeUseCase := usecase.NewListSafesUseCase(SafeRepository)
	ListSafeController := controller.NewListSafesController(ListSafeUseCase)

	return ListSafeController
}
