package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeLowTodayTransaction() controller.GetLowTodayController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	MakeLowTodayTransactionController := controller.NewLGetLowTodayController(TransactionRepository)

	return MakeLowTodayTransactionController
}
