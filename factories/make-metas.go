package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/meta"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/meta"
)

func MakeMetas() controller.GetMetasController {
	MetaRepository := repository.NewMetaRepository(db.RunDB())
	GetMetasUseCase := usecase.NewGetMetasUseCase(MetaRepository)
	GetMetasController := controller.NewGetMetasController(GetMetasUseCase)

	return GetMetasController
}
