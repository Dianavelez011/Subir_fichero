package excel

import (
	"fmt"
	"github.com/gin/internal/services/file"

	// "fmt"

	"github.com/xuri/excelize/v2"
)

func (s Service)SaveContent(fileName string)error {
	//open file
	//get content
	//translate content
	//save content

	excel,err := excelize.OpenFile(s.FileLocation+fileName)

	if err != nil{
		return fmt.Errorf("error open excel file: %s",err.Error())
	}


	//Obtener una lista de todas las hojas del archivo
	sheets := excel.GetSheetList()
	
	for _,sheetName := range sheets{
		//Obtener las filas de la hoja actual del archivo
		rows,err := excel.GetRows(sheetName)
		if err != nil {
			return fmt.Errorf("error to read sheet %s row: %w",sheetName,err)
		}

		rowsInterface := file.ToInterfaceSlice(rows);

		
		fmt.Println(rows)
		//Insertar en la base de datos
		if err := s.Repo.CopyFrom(s.Columns,rowsInterface,s.TableName); err != nil{
			return err
		}

	}
	return nil
}