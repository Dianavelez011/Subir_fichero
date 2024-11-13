package excel

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

)

var (
	mapSizeMainFile = make(map[string]map[string]int)
)

func (s Service) Join(partFile *multipart.FileHeader, sizeMain int,fileName string) (bool, error) {
	//create or open main file
	//open part of the file
	//read chunk content in buffer
	//write chunk in main file
	fmt.Printf("tamaño parte archivo %d, tamaño archivo principal %d",partFile.Size,sizeMain)

	mainFilePath := s.FileLocation + fileName
	valueMap,exists := mapSizeMainFile[fileName];

	if !exists{
		mapSizeMainFile[fileName] = map[string]int{
			"sizeMainFile": sizeMain,
			"sizeCompleted": 0,
		}

		valueMap = mapSizeMainFile[fileName]
	}

	
	//Initial value to sizeMainFile
	// if sizeMainFile == 0 {
	// 	sizeMainFile = sizeMain
	// }
	//opoen file if not exist it will create new file
	file, err := os.OpenFile(mainFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, fmt.Errorf("unable to open file for appending in Join: %s", err.Error())
	}
	defer file.Close()

	chunkFile, err := partFile.Open()
	if err != nil {
		return false, fmt.Errorf("unabled to open file chunk:%s", err.Error())
	}

	defer chunkFile.Close()

	// read the chunk content in bytes buffer
	buffer := make([]byte, partFile.Size)
	for {
		n, err := chunkFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return false, fmt.Errorf("failed to read chunk file: %s", err.Error())
		}
		// Escribe sólo los bytes leídos, hasta n
		if _, err := file.Write(buffer[:n]); err != nil {
			return false, fmt.Errorf("failed to write chunk to file: %s", err.Error())
		}
	}

	// sizeComplete += int(partFile.Size)
	valueMap["sizeCompleted"] += int(partFile.Size)

	// if sizeComplete == sizeMainFile {
	// 	sizeMainFile = 0
	// 	sizeComplete = 0
	// 	return true, nil
	// }
	if valueMap["sizeCompleted"] == valueMap["sizeMainFile"] {
		delete(mapSizeMainFile,fileName)
		return true, nil
	}

	return false, nil

}
