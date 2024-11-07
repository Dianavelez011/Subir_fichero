package excel

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (h Handler) Create(ctx *gin.Context, file *multipart.FileHeader) {
	//Traducir el request
	//Validacion
	//Consumir un servicio
	//Traducir el response




	


	
	//--------------------------------------------

	err := h.FileService.Create(ctx,file)
	if err != nil{
		fmt.Printf("error creating file: %s",err.Error())
		return
	}

	//-----------------------------------------------------------

	ctx.JSON(200, gin.H{"message": "file upload success!"})

}
