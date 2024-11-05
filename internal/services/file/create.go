package file

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s Service)Create(ctx *gin.Context, file *multipart.FileHeader)error {
	//validate type file
	//clean file if type file is .txt
	//save in database
	//response success

	typeFile, err := TypeFile(file.Filename)

	if err != nil {
		fmt.Printf("error: %s",err.Error());
		return fmt.Errorf("error unsupported media type:%s",err.Error())
	}


	switch typeFile {
	case ".xls", ".xlsx":
		if err := s.ExcelService.Create(ctx,file); err != nil{
			return err
		}
		
	case ".txt":
		if err:=s.TxtService.Create(ctx,file); err != nil{
			return err
		}

	default:
		return fmt.Errorf("unsuported file type: %s",typeFile)
	}

	return nil
}