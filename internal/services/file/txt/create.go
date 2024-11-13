package txt

import (
	"mime/multipart"
	// "sync"

	"github.com/gin-gonic/gin"
)

func (s Service) Create(ctx *gin.Context, file *multipart.FileHeader, sizeMainFile int,channel chan <- map[string]interface{}) {
	//validate
	// join file
	// save file
	//save content
	cleanFileName := s.CleanName(file.Filename)
	mainFilePath := s.FileLocation + cleanFileName
	response := map[string]interface{}{
		"error": nil,
		"error_message": "Internal server error",
		"code": 500,
		"upload_file": false,
		"path_file":"",
	}

	// var wg sync.WaitGroup
	// channel := make(chan error)

	// defer wg.Done()


	if err := s.ValidateChunk(file.Filename, "parte"); err != nil {
		response["error"] = err
		response["error_message"] = "File stream submission failed";
		response["code"] = 400
		channel <- response
		s.Delete(mainFilePath)
		return
	}
	//clean file name
	
	
	buitFile, err := s.Join(file, sizeMainFile, cleanFileName)

	if err != nil {
		response["error"] = err
		s.Delete(mainFilePath)
		channel <- response
		return
	}
	

	if !buitFile {
		channel <- response
		return
	}else{
		response["upload_file"] = true
		response["path_file"] = mainFilePath
		channel <- response
		return
	}


	// fmt.Println("Linea 55")
	// ctx.JSON(200, gin.H{"message": "txt file upload success!"})


	// err = s.SaveContent(mainFilePath)

	// if err != nil{
	// 	response["error"] = err
	// 	s.Delete(mainFilePath)
	// 	channel <- response
	// 	return
	// }else{
	// 	s.Delete(mainFilePath)
	// 	channel <- response
	// }


	// channel <- nil 


	// wg.Add(1)
	// go s.SaveContent(&wg, channel, mainFilePath)

	// select {
	// case <-ctx.Done():
	// 	s.Delete(mainFilePath)
	// 	return ctx.Err()

	// case err := <-channel:
	// 	if err != nil {
	// 		return fmt.Errorf("could save txt content in SaveContent:%s", err.Error())
	// 	}else{
	// 		s.Delete(mainFilePath)
	// 	}

	// }

	// wg.Wait()
	// close(channel)

	// if err := ctx.SaveUploadedFile(file, s.FileLocation+file.Filename); err != nil {
	// 	return fmt.Errorf("could not save file:%s",err.Error())
	// }

	// fmt.Println(file.Filename)

	// return nil
}
