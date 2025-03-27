package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/supplier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeSupplier() controller.SupplierController {
	SupplierRepository := repository.NewSupplierRepository(db.RunDB())
	CreateSupplierController := controller.NewSupplierController(*SupplierRepository)

	return CreateSupplierController
}
