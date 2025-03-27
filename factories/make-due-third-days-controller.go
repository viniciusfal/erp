package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeDueLastThirdDays() controller.GetDueLastThirdyDaysController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	MakeDueLastThirdDaysController := controller.NewLGetDueLastThirdyDaysController(TransactionRepository)

	return MakeDueLastThirdDaysController
}
