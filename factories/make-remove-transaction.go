package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func MakeRemoveTransaction() controller.RemoveTransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	RemoveTransactionUseCase := usecase.NewRemoveTransactionUseCase(TransactionRepository)
	RemoveTransactionController := controller.NewRemoveTransactionController(RemoveTransactionUseCase)

	return RemoveTransactionController
}
