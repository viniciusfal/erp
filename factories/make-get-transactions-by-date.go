package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func MakeGetTransactionByDate() controller.GetTransactionByDateController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	GetTransactionByDateUseCase := usecase.NewGetTransactionByDateUseCase(TransactionRepository)
	GetTransactionByDateController := controller.NewGetTransactionByDateController(GetTransactionByDateUseCase)

	return GetTransactionByDateController
}
