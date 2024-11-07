package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

func MakeAnalysesTransactionsByMonth() controller.AnalysesTransactionController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	AnalysesTransactionUseCase := usecase.NewAnalysesTransactionUseCase(TransactionRepository)
	AnalysesTransactionController := controller.NewAnalysesTransactionController(AnalysesTransactionUseCase)

	return AnalysesTransactionController
}
