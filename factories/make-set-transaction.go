package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

func MakeSetTransaction() controller.SetTransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	SetTransactionUseCase := usecase.NewSetTransactionUseCase(TransactionRepository)
	SetTransactionController := controller.NewSetTransactionController(SetTransactionUseCase)

	return SetTransactionController
}
