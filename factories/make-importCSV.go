package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

func MakeImportTransactionsCSV() controller.ImportCSVController {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	ImportCSVUseCase := usecase.NewImportCSCVUseCase(TransactionRepository)
	ImportCSVController := controller.NewImportSCVController(ImportCSVUseCase)

	return ImportCSVController
}
