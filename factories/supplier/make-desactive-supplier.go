package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/supplier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeDesactiveSupplier() controller.DesactiveSupplierController {
	SupplierRepository := repository.NewSupplierRepository(db.RunDB())
	DesactiveSupplierController := controller.NewDesactiveSupplierController(*SupplierRepository)

	return DesactiveSupplierController
}
