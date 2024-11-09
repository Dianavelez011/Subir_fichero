package txt

import (
	"context"
	"mime/multipart"
	"sync"
)

func (s Service) Create(ctx context.Context, file *multipart.FileHeader, sizeMainFile int,channel chan map[string]interface{},wg *sync.WaitGroup) {
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
	}

	// var wg sync.WaitGroup
	// channel := make(chan error)

	defer wg.Done()


	if err := s.ValidateChunk(file.Filename, "parte"); err != nil {
		response["error"] = err
		response["error_message"] = "File stream submission failed";
		response["code"] = 400
		channel <- response
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
	}

	channel <- response

	err = s.SaveContent(mainFilePath)

	if err != nil{
		response["error"] = err
		s.Delete(mainFilePath)
		channel <- response
		return
	}else{
		s.Delete(mainFilePath)
	}


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
