package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

func MakeGetTransactionByDate() controller.GetTransactionByDateController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	GetTransactionByDateUseCase := usecase.NewGetTransactionByDateUseCase(TransactionRepository)
	GetTransactionByDateController := controller.NewGetTransactionByDateController(GetTransactionByDateUseCase)

	return GetTransactionByDateController
}
