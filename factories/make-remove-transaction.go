package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

func MakeRemoveTransaction() controller.RemoveTransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	RemoveTransactionUseCase := usecase.NewRemoveTransactionUseCase(TransactionRepository)
	RemoveTransactionController := controller.NewRemoveTransactionController(RemoveTransactionUseCase)

	return RemoveTransactionController
}
