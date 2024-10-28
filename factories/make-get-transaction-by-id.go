package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

func MakeGetTransactionById() controller.GetTransactionByIdController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	GeTransactionByIdUseCase := usecase.NewGetTransactionByIdUseCase(TransactionRepository)
	GetTransactionByIdController := controller.NewGetTransactionByIdController(GeTransactionByIdUseCase)

	return GetTransactionByIdController
}
