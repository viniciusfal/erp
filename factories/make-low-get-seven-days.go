package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeLowGetSevenDays() controller.GetLowLastSevenDays {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	MakeLowGetSevenDaysController := controller.NewLowGetLastSevenDays(TransactionRepository)

	return MakeLowGetSevenDaysController
}
