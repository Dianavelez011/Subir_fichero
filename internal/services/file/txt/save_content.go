package txt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"

	// "github.com/jackc/pgx/v5/pgxpool"
	// "sync"
	// "github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/text/transform"
)

func (s Service) SaveContent(mainFilePath string) error {
	//open file
	//read file
	//decode content
	//save in database

	//mainfile path

	openedFile, err := os.Open(mainFilePath)
	channel := make(chan map[string]interface{})
	var wg sync.WaitGroup
	// var db_connection *pgxpool.Pool

	if err != nil {
		return fmt.Errorf("could not open file: %s", err.Error())
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
			fmt.Println("Linea 40")
			decodeChunk, _, decodeErr := transform.Bytes(s.Decoder, buffer[:n])
			if decodeErr != nil {
				return fmt.Errorf("decoding error :%s",decodeErr.Error())
			}
			data := s.ProccessTextToSlice(string(decodeChunk))
			dataFilter := s.FilterRows(data)

			rowsInterface := s.ToInterfaceSlice(dataFilter)
			// (string(decodeChunk), 53, "UNIDAD VICTIMAS")
			wg.Add(1)
			go s.Repo.CopyFrom(s.Columns, rowsInterface, s.TableName, channel, &wg)

			if err := <-channel; err["error"] != nil {
				return fmt.Errorf("error could not execute copyfrom in SaveContent:%s", err["error"].(error).Error())
			}
			// else{
			// 	result := <-channel
			// 	db_connection = result["db_connection"].(*pgxpool.Pool)
				
			// }
			

			// fmt.Println(data)
		}

		if err != nil {
			fmt.Println("Linea 54")

			if err == io.EOF {
				fmt.Println("Linea 68")
				break

			}
			fmt.Println("Linea 79")

			return fmt.Errorf("could not read file in SaveContent:%s", err.Error())
		}
		fmt.Println("Linea 61")

	}
	wg.Wait()
	// db_connection.Close()
	close(channel)
	fmt.Println("rows copied succesfully")

	return nil

}
