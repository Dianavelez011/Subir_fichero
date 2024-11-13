package excel

import (
	"fmt"
	"os"
)

func (s Service) DeleteFolder() error{
	if err := os.RemoveAll(s.FileLocation); err != nil{
		return fmt.Errorf("could not delete folder in DeleteFolder: %s",err.Error())
	}
	return nil
}