package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeMarkLowTransactionController() controller.MarkLowTransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	MarkLowTransactionController := controller.NewMarkLowTransactionController(TransactionRepository)

	return MarkLowTransactionController
}
