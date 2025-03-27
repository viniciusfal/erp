package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/supplier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeListSupplierWithTransactions() controller.ListWithTransactionController {
	SupplierRepository := repository.NewSupplierRepository(db.RunDB())
	ListWithSupplierController := controller.NewSupplierWithTransactionsController(*SupplierRepository)

	return ListWithSupplierController
}
