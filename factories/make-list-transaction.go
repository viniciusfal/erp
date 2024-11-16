package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

func MakeListTransactions() controller.ListTransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	ListTransactionUseCase := usecase.NewListTransactionUseCase(TransactionRepository)
	ListTransactionController := controller.NewListTransactionController(ListTransactionUseCase)

	return ListTransactionController
}
