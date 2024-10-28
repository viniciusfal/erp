package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/safe"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/safe"
)

func MakeSetSafe() controller.SetSafeController {
	SafeRepository := repository.NewSafeRepository(db.RunDB())
	SetSafeUseCase := usecase.NewSetSafeUseCase(SafeRepository)
	SetSafeController := controller.NewSetSafeController(SetSafeUseCase)

	return SetSafeController
}
