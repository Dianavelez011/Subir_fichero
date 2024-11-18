package txt

import (
"os"
"fmt"
)

func (s *Service) Delete(fileName string) error{
	s.UploadFile = true;
	mainFilePath := s.FileLocation + fileName
	err := os.Remove(mainFilePath)

	if err != nil{
		return fmt.Errorf("error deleted file in Dete: %s",err.Error())
	}
	return nil
}