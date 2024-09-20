package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func MakeTransactions() controller.TransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	TransactionUseCase := usecase.NewTransactionUseCase(TransactionRepository)
	TransactionController := controller.NewTransactionController(TransactionUseCase)

	return TransactionController
}
