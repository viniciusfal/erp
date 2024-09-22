package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func MakeSetTransaction() controller.SetTransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	SetTransactionUseCase := usecase.NewSetTransactionUseCase(TransactionRepository)
	SetTransactionController := controller.NewSetTransactionController(SetTransactionUseCase)

	return SetTransactionController
}
