package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeCreateLowTo0Controller() controller.CreateLowTo0TransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	CreateLowTo0TransactionController := controller.NewCreateLowTo0TransactionController(TransactionRepository)

	return CreateLowTo0TransactionController
}
