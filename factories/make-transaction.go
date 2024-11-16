package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

func MakeTransactions() controller.TransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	TransactionUseCase := usecase.NewTransactionUseCase(TransactionRepository)
	TransactionController := controller.NewTransactionController(TransactionUseCase)

	return TransactionController
}
