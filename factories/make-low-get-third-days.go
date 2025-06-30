package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeLowGetThirdDays() controller.GetLowThirdyDays {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	GetLowThirdyDaysController := controller.NewLGetLowThirdyDays(TransactionRepository)

	return GetLowThirdyDaysController
}
