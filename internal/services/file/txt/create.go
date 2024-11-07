package txt

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s Service)Create(ctx *gin.Context,file *multipart.FileHeader,sizeMainFile int )error {
	//validate
	// join file
	// save file
	//save content
	fmt.Println(sizeMainFile)

	buitFile,err := s.Join(ctx,file,sizeMainFile)

	if err != nil{
		return fmt.Errorf("failed join file:%s",err.Error())
	}

	if buitFile {
		if err := s.SaveContent(); err != nil{
			return fmt.Errorf("could save txt content in SaveContent:%s",err.Error())
		}
	}else{
		return nil
	}


	// if err := ctx.SaveUploadedFile(file, s.FileLocation+file.Filename); err != nil {
	// 	return fmt.Errorf("could not save file:%s",err.Error())
	// }

	// fmt.Println(file.Filename)

	

	return nil
}