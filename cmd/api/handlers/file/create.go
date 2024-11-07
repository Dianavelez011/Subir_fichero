package file

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin/cmd/api/handlers/file/excel"
	"github.com/gin/cmd/api/handlers/file/txt"
)

func (h Handler) Create(ctx *gin.Context) {
	//Traducir el request
	//Validacion
	//Consumir un servicio
	//Traducir el response

	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Printf("error:%s", err.Error())
		ctx.JSON(400, gin.H{"error": "file not found"})
		return
	}

	typeFile, err := h.TypeFile(file.Filename)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		ctx.JSON(415, gin.H{"error": "unsupported media type"})
		return
	}

	switch typeFile {
	case ".xls", ".xlsx":
		excelHandler := excel.Handler{
			FileService: h.ExcelService,
		}
		excelHandler.Create(ctx, file)

	case ".txt":
		txtHandler := txt.Handler{
			FileService: h.TxtService,
		}
		txtHandler.Create(ctx, file)

	}

}
