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

	contextTimeout, _ := context.WithTimeout(context.Background(), 7*time.Minute)
	channel := make(chan error)
	var wg sync.WaitGroup

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
		ctx.JSON(408, gin.H{"error": "timeout waiting for the next fragment request canceled"})
		return
	case err := <-channel:
		if err != nil {
			// if errors.Is(err, context.DeadlineExceeded) {
			// 	ctx.JSON(408, gin.H{"error": "timeout waiting for the next fragment request canceled"})
			// } else {
				fmt.Printf("filed to create file: %s", err.Error())
				ctx.JSON(500, gin.H{"error": "failed to create file"})
			// }
			return
		}

	}

		wg.Wait()
		close(channel)

	//-----------------------------------------------------------

	ctx.JSON(200, gin.H{"message": "txt file upload success!"})

}
