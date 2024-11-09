package excel

import (
	"mime/multipart"
	"sync"

	"github.com/gin-gonic/gin"
)

type FileService interface {
	Create(ctx *gin.Context, file *multipart.FileHeader)error
}

type FileRepository interface {
	// InsertOrUpdate(query string, values []interface{}) error
	CopyFrom(columns []string, values [][]interface{},tableName string,channel chan map[string]interface{},wg *sync.WaitGroup)
}