package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func MakeListTransactions() controller.ListTransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	ListTransactionUseCase := usecase.NewListTransactionUseCase(TransactionRepository)
	ListTransactionController := controller.NewListTransactionController(ListTransactionUseCase)

	return ListTransactionController
}
