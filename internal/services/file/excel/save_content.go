package excel

import (
	"fmt"
	// "fmt"

	"github.com/xuri/excelize/v2"
)

func (s Service)SaveContent(fileName string)error {
	//open file
	//get content
	//translate content
	//save content

	fmt.Println("SaveContent line 18")

	excel,err := excelize.OpenFile(s.FileLocation+fileName)
	fmt.Println("SaveContent line 22")


	if err != nil{
		return fmt.Errorf("error open excel file: %s",err.Error())
	}

	// defer excel.Close()


	//Obtener una lista de todas las hojas del archivo
	sheets := excel.GetSheetList()
	fmt.Println("SaveContent line 33")
	
	for _,sheetName := range sheets{
		fmt.Println(sheetName)
		//Obtener las filas de la hoja actual del archivo
		rows,err := excel.GetRows(sheetName)
		if err != nil {
			return fmt.Errorf("error to read sheet %s row: %w",sheetName,err)
		}
		fmt.Println("SaveContent line 41")
		//Convert slice[][]string to slice [][]interface{}
		rowsInterface := s.ToInterfaceSlice(rows);

		//Insertar en la base de datos
		if err := s.Repo.CopyFrom(s.Columns,rowsInterface,s.TableName); err != nil{
			return fmt.Errorf("could not execute copyFrom:%s",err.Error())
		}

	}
	return nil
}