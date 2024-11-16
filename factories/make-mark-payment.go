package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

func MakeMarkPayment() controller.MarkPaymentController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	MarkPaymentUseCase := usecase.NewMarkPaymentUseCase(TransactionRepository)
	MarkPaymentController := controller.NewMarkPaymentUseController(MarkPaymentUseCase)

	return MarkPaymentController
}
