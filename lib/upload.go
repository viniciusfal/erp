package lib

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)

	dst := filepath.Join("../uploads", file.Filename)

	ctx.SaveUploadedFile(file, dst)

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
