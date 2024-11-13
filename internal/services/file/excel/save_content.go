package excel

import (
	"fmt"
	"sync"

	// "fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xuri/excelize/v2"
)

func (s Service)SaveContent(mainFilePath string)error {
	//open file
	//get content
	//translate content
	//save content

	fmt.Println("SaveContent line 18")

	excel,err := excelize.OpenFile(mainFilePath)
	channel  := make(chan map[string]interface{})
	var wg sync.WaitGroup
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
		wg.Add(1)
		go s.Repo.CopyFrom(s.Columns,rowsInterface,s.TableName,channel,&wg)
		if err:= <-channel; err["error"].(error) != nil{
			return fmt.Errorf("error could not execute copyfrom in SaveContent:%s",err["error"].(error).Error())
		}
		wg.Wait()
		
	}
	dbConnection := <-channel
	close(channel)

	defer dbConnection["db_connection"].(*pgxpool.Pool).Close()
	return nil
}