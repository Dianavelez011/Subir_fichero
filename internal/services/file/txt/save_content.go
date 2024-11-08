package txt

import (
	"bufio"
	"fmt"
	"os"
	// "sync"

	"golang.org/x/text/transform"
)

func (s Service) SaveContent(mainFilePath string)error{
	//open file
	//read file
	//decode content
	//save in database

	//mainfile path

	openedFile, err := os.Open(mainFilePath)

	if err != nil {
		return fmt.Errorf("could not open file: %s",err.Error())
	}

	//Cerrar el archivo
	defer openedFile.Close()
	//close waitGroup
	// defer wg.Done()

	reader := bufio.NewReader(openedFile)
	//Tamaño buffer 125mb
	buffer := make([]byte, 125*1024*1024)

	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			decodeChunk, _, decodeErr := transform.Bytes(s.Decoder, buffer[:n])
			if decodeErr != nil {
				return fmt.Errorf("decoding error :%s",decodeErr.Error())
			}
			data := s.ProccessTextToSlice(string(decodeChunk))

			rowsInterface := s.ToInterfaceSlice(data)
			// (string(decodeChunk), 53, "UNIDAD VICTIMAS")
			s.Repo.CopyFrom(s.Columns,rowsInterface,s.TableName)
			// fmt.Println(data)
		}

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			 return fmt.Errorf("could not read file in SaveContent:%s",err.Error())
		}
	}
	fmt.Println("rows copied succesfully")

	return nil
}