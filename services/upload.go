package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadSingleFile(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")

	log.Println(file.Filename)

	ctx.SaveUploadedFile(file, "./uploads")

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

}
