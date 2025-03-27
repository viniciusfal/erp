package factories

import (
	"github.com/viniciusfal/erp/db"
	controller "github.com/viniciusfal/erp/http/controller/supplier"
	"github.com/viniciusfal/erp/infra/repository"
)

func MakeSetSupplier() controller.SetSupplierController {
	SupplierRepository := repository.NewSupplierRepository(db.RunDB())
	SetSupplierController := controller.NewSetSupplierController(*SupplierRepository)

	return SetSupplierController
}
