package services

import (
	"fmt"
	"mime/multipart"

	"github.com/xuri/excelize/v2"
)

func CleanFile(file *multipart.FileHeader)error {
	openedFile,err := file.Open();

	if err != nil{
		return fmt.Errorf("could not open file:%w",err.Error())
	}

	//Cerrar el archivo
	defer openedFile.Close()

	excel,err := excelize.OpenReader(openedFile)

	if err != nil{
		return fmt.Errorf("error open excel file: %w",err.Error())
	}


	//Obtener una lista de todas las hojas del archivo
	sheets := excel.GetSheetList()
	
	for _,sheetName := range sheets{
		//Obtener las filas de la hoja actual del archivo
		rows,err := excel.GetRows(sheetName)
		if err != nil {
			return fmt.Errorf("error to read sheet %s row: %w",sheetName,err)
		}
		
		fmt.Println(rows)
		//Insertar en la base de datos
		db.Insert(rows)

	}
}

func CleanTxtFile(file *multipart.FileHeader){

}