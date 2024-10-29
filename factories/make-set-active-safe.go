package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/safe"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/safe"
)

func MakeSetActive() controller.SetActiveController {
	SafeRepository := repository.NewSafeRepository(db.RunDB())
	SetActiveUseCase := usecase.NewSetActiveUseCase(SafeRepository)
	SetActiveController := controller.NewSetActiveController(SetActiveUseCase)

	return SetActiveController
}
