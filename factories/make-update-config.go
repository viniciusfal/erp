package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/config"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/config"
)

func MakeUpdateConfig() controller.UpdateConfigController {
	configRepository := repository.NewConfigRepository(db.RunDB())
	updateConfigUseCase := usecase.NewUpdateConfigUseCase(configRepository)
	updateConfigController := controller.NewUpdateConfigController(updateConfigUseCase)

	return updateConfigController
} 