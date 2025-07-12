package factories

import (
	file "github.com/viniciusfal/erp/http/controller/file"
)

func MakeUploadController() file.UploadController {
	return file.NewUploadController()
} 