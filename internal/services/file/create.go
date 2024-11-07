package file

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s Service)Create(ctx *gin.Context, file *multipart.FileHeader)error {
	//validate type file
	//clean file if type file is .txt
	//save in database
	//response success

	return nil
}