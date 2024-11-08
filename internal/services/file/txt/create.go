package txt

import (
	"context"
	"fmt"
	"mime/multipart"
	"sync"
)

func (s Service) Create(ctx context.Context, file *multipart.FileHeader, sizeMainFile int,channel chan error,wg *sync.WaitGroup) {
	//validate
	// join file
	// save file
	//save content
	fmt.Println(sizeMainFile)

	// var wg sync.WaitGroup
	// channel := make(chan error)

	defer wg.Done()


	if err := s.ValidateChunk(file.Filename, "parte"); err != nil {
		channel <- err
		return
	}
	//clean file name
	
	cleanFileName := s.CleanName(file.Filename)
	mainFilePath := s.FileLocation + cleanFileName
	buitFile, err := s.Join(file, sizeMainFile, cleanFileName)

	if err != nil {
		channel <- fmt.Errorf("failed join file:%s", err.Error())
		return
	}

	if !buitFile {
		channel <- nil
		return
	}

	err = s.SaveContent(mainFilePath)

	if err != nil{
		channel <- err 
		return
	}

	channel <- nil 


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
