package file

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h Handler) Create(ctx *gin.Context) {
	//Traducir el request
	//Validacion
	//Consumir un servicio
	//Traducir el response


	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Printf("error:%s",err.Error())
		ctx.JSON(400, gin.H{"error": "file not found"})
		return
	}

	
	//--------------------------------------------

	err = h.FileService.Create(ctx,file)
	if err != nil{
		fmt.Printf("error creating file: %s",err.Error())
		return
	}

	//-----------------------------------------------------------

	ctx.JSON(200, gin.H{"message": "file upload success!"})

}
