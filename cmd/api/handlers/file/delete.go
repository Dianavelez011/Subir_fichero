package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h Handler) Delete(ctx *gin.Context) {
	fileName := ctx.Param("fileName")
	cleanName := h.CleanName(fileName)
	err := h.TxtService.Delete(cleanName)

	if err != nil {
		fmt.Printf("could not delete folder: %s",err.Error())
		ctx.JSON(500, gin.H{"message": "Internal server error"})
		return
	}

	ctx.JSON(200, gin.H{"message": "txt file deleted success!"})
}