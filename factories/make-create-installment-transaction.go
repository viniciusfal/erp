package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeCreateInstallmentTransactionController() controller.CreateInstallmentTransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	CreateInstallmentTransactionController := controller.NewCreateInstallmentTransactionController(TransactionRepository)

	return CreateInstallmentTransactionController
}
