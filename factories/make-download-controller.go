package factories

import (
	file "github.com/viniciusfal/erp/http/controller/file"
)

func MakeDownloadController() file.DownloadController {
	return file.NewDownloadController()
} 