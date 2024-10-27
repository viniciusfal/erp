package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func MakeMeta() controller.CreateMetaController {
	MetaRepository := repository.NewMetaRepository(db.RunDB())
	CreateMetaUseCase := usecase.NewMetaUseCase(MetaRepository)
	MetaController := controller.NewCeateMetaController(CreateMetaUseCase)

	return MetaController
}
