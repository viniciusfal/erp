package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/supplier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeListSupplier() controller.ListAllSupplierController {
	SupplierRepository := repository.NewSupplierRepository(db.RunDB())
	ListSupplierController := controller.NewListAllSupplierController(*SupplierRepository)

	return ListSupplierController
}
