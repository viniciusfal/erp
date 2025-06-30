package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/supplier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeFindOneSupplier() controller.FindOneSupplierController {
	SupplierRepository := repository.NewSupplierRepository(db.RunDB())
	FindOneSupplierController := controller.NewFindOneSupplierController(*SupplierRepository)

	return FindOneSupplierController
}
