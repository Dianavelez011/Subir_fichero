package file

import (
	"github.com/gin-gonic/gin"
)

(h Handler)func CreateFile(ctx *gin.Context) {
	//Traducir el request
	//Validacion
	//Consumir un servicio
	//Traducir el response

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//--------------------------------------------


	//-----------------------------------------------------------

	ctx.JSON(200, gin.H{"message": "file upload success!"})

}

