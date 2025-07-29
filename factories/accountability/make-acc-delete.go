package factories

import (
	"github.com/viniciusfal/erp/db"
	"github.com/viniciusfal/erp/infra/repository"
	"github.com/viniciusfal/erp/http/controller/accountability"
)

func MakeAccDelete() controller.DeleteAccController {
	accRepo := repository.NewAccountabilityRepository(db.RunDB())
	return controller.NewDeleteAccController(accRepo)
} 