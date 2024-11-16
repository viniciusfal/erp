package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/meta"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/meta"
)

func MAkeSetMeta() controller.SetMetaController {
	MetaRepository := repository.NewMetaRepository(db.RunDB())
	SetMetaUseCase := usecase.NewSetMetaUseCase(MetaRepository)
	SetMetaController := controller.NewSetMetaController(SetMetaUseCase)

	return SetMetaController
}
