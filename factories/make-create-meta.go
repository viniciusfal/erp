package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/meta"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/meta"
)

func MakeMeta() controller.CreateMetaController {
	MetaRepository := repository.NewMetaRepository(db.RunDB())
	CreateMetaUseCase := usecase.NewMetaUseCase(MetaRepository)
	MetaController := controller.NewCeateMetaController(CreateMetaUseCase)

	return MetaController
}
