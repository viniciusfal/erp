package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/config"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/config"
)

func MakeGetConfig() controller.GetConfigController {
	configRepository := repository.NewConfigRepository(db.RunDB())
	getConfigUseCase := usecase.NewGetConfigUseCase(configRepository)
	getConfigController := controller.NewGetConfigController(getConfigUseCase)

	return getConfigController
} 