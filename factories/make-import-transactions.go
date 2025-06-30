package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/transaction"
	"github.com/viniciusfal/erp/infra/repository"
	usecase "github.com/viniciusfal/erp/infra/usecase/transaction"
)

func MakeImportTransactions() controller.TransactionFromExcel {
	TransactionRepository := repository.NewTransactionRepository(db.RunDB())
	ImportTransactionUseCase := usecase.NewTransactionImport(TransactionRepository)
	ImportTransactionController := controller.NewTransactionFromExcel(ImportTransactionUseCase)

	return ImportTransactionController
}
