package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (h Handler) Delete(ctx *gin.Context) {
	err := h.TxtService.DeleteFolder()

	if err != nil {
		fmt.Printf("could not delete folder: %s",err.Error())
	}

	ctx.JSON(200, gin.H{"message": "txt file upload success!"})
}