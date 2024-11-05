package excel

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s Service)Create(ctx *gin.Context, file *multipart.FileHeader)error {
	if err := ctx.SaveUploadedFile(file, s.FileLocation + file.Filename); err != nil {
		return fmt.Errorf("could not save file:%s",err.Error())
	}

	if err := s.SaveContent(file.Filename); err != nil{
		return fmt.Errorf("could save excel content in SaveContent:%s",err.Error())
	}

	


	return nil
}