package services

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func CreateFile(ctx *gin.Context, file *multipart.FileHeader) {

	if err := ctx.SaveUploadedFile(file, "./uploads/"+file.Filename); err != nil {
		fmt.Printf("error:%w",err.Error())
		ctx.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	typeFile, err := TypeFile(file.Filename)

	if err != nil {
		fmt.Printf("error: %w",err.Error());
		ctx.JSON(415, gin.H{"error": "unsupported media type"})
		return
	}

	switch typeFile {
	case ".xls", ".xlsx":

	}

	if mimetype == "text/plain; charset=utf-8" {
		files.GetCsvData(file.Filename)
	} else {
		files.GetExcelData(file.Filename)
	}

}