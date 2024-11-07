package txt

import (
	"fmt"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) Create(ctx *gin.Context,file *multipart.FileHeader) {
	//Traducir el request
	//Validacion
	//Consumir un servicio
	//Traducir el response


	sizeMainFileStr:= ctx.PostForm("sizeMainFile")
	sizeMainFile,err := strconv.Atoi(sizeMainFileStr) 

	if err != nil {
		fmt.Printf("sizeMainFile is not a valid integer in Create: %s", err.Error())
		ctx.JSON(400,gin.H{"error":"Invalid data format. Please ensure the data is formatted correctly."})
	}
	
	

	
	//--------------------------------------------

	err = h.FileService.Create(ctx,file,sizeMainFile)
	// err = h.FileService.Create(ctx,file)
	if err != nil{
		fmt.Printf("error creating file: %s",err.Error())
		return
	}

	//-----------------------------------------------------------

	ctx.JSON(200, gin.H{"message": "txt file upload success!"})

}
