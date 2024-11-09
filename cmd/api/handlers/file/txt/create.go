package txt

import (
	"context"
	// "errors"
	"fmt"
	"mime/multipart"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func (h Handler) Create(ctx *gin.Context, file *multipart.FileHeader) {
	//Traducir el request
	//Validacion
	//Consumir un servicio
	//Traducir el response

	contextTimeout, cancel := context.WithTimeout(context.Background(), 7*time.Minute)
	channel := make(chan map[string]interface{})
	var wg sync.WaitGroup

	defer cancel()

	sizeMainFileStr := ctx.PostForm("sizeMainFile")
	sizeMainFile, err := strconv.Atoi(sizeMainFileStr)

	if err != nil {
		fmt.Printf("sizeMainFile is not a valid integer in Create: %s", err.Error())
		ctx.JSON(400, gin.H{"error": "Invalid data format. Please ensure the data is formatted correctly."})
		return
	}

	//--------------------------------------------
	wg.Add(1)
	go h.FileService.Create(contextTimeout, file, sizeMainFile, channel, &wg)

	select {
	case <-ctx.Done():
		if err:= h.FileService.DeleteFolder(); err!=nil{
			fmt.Printf("failed delete folder:%s",err.Error())
			return
		}
		ctx.JSON(408, gin.H{"error": "timeout waiting for the next fragment request canceled"})
		return
	case response := <-channel:
		if err := response["error"]; err != nil {
			fmt.Printf("filed to create file: %s", err.(error).Error())
			ctx.JSON(response["code"].(int), gin.H{"error": response["error_message"]})
			return
		}else{
			ctx.JSON(200, gin.H{"message": "txt file upload success!"})
		}

	}

	wg.Wait()
	close(channel)

	//-----------------------------------------------------------

}
