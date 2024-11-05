package txt

import (
	"bufio"
	"fmt"
	"os"
	"golang.org/x/text/transform"
	"github.com/gin/internal/services/file"
)

func (s Service) SaveContent(fileName string) error{
	//open file
	//read file
	//decode content


	openedFile, err := os.Open(s.FileLocation+fileName)

	if err != nil {
		return fmt.Errorf("could not open file: %s",err.Error())
	}

	//Cerrar el archivo
	defer openedFile.Close()

	reader := bufio.NewReader(openedFile)
	//Tamaño buffer 125mb
	buffer := make([]byte, 125*1024*1024)

	// decoder := charmap.ISO8859_1.NewDecoder()
	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			decodeChunk, _, decodeErr := transform.Bytes(s.Decoder, buffer[:n])
			if decodeErr != nil {
				return fmt.Errorf("decoding error :%w",decodeErr.Error())
			}
			data := s.ProccessTextToSlice(string(decodeChunk))

			rowsInterface := file.ToInterfaceSlice(data)
			// (string(decodeChunk), 53, "UNIDAD VICTIMAS")
			s.Repo.CopyFrom(s.Columns,rowsInterface,s.TableName)
			fmt.Println(data)
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