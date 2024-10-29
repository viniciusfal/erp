package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/meta"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/meta"
)

func MakeGetMetaByMonth() controller.GetMetaByMonthController {
	MetaRepository := repository.NewMetaRepository(db.RunDB())
	GetMetasByMonthUseCase := usecase.NewMetasByMonthUseCase(MetaRepository)
	GetMetaByMonthController := controller.NewGetMetaByMonthUseCase(GetMetasByMonthUseCase)

	return GetMetaByMonthController
}
