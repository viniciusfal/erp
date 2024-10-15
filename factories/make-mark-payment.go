package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/http/controller"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/infra/usecase"
)

func MakeMarkPayment() controller.MarkPaymentController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	MarkPaymentUseCase := usecase.NewMarkPaymentUseCase(TransactionRepository)
	MarkPaymentController := controller.NewMarkPaymentUseController(MarkPaymentUseCase)

	return MarkPaymentController
}
