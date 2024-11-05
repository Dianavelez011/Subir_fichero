package file

import (
	"fmt"
	"path/filepath"

)

//Almacenar en un map los tipo de archivos permitidos
var allowedMimeTypes = map[string]bool{
	".xlsx":true, // Excel .xlsx
	".xls":true, // Excel .xls
	".txt":true, //Txt
}


func (h Handler)TypeFile(fileName string) (string,error){
	//Fallar si el tipo de archivo no esta permitido
	ext := filepath.Ext(fileName)
	if !allowedMimeTypes[ext] {
		return "",fmt.Errorf("type file %s dont allowed:",ext)
	}

	return ext,nil
}