package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func MakeGetTransactionById() controller.GetTransactionByIdController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	GeTransactionByIdUseCase := usecase.NewGetTransactionByIdUseCase(TransactionRepository)
	GetTransactionByIdController := controller.NewGetTransactionByIdController(GeTransactionByIdUseCase)

	return GetTransactionByIdController
}
