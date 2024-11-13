package txt

import (
	"mime/multipart"
	"sync"

	"github.com/gin-gonic/gin"
)

type FileService interface {
	Create(ctx *gin.Context, file *multipart.FileHeader, sizeMainFile int,channel chan <- map[string]interface{})
	DeleteFolder() error
	SaveContent(mainFilePath string) error
	Delete(mainFilePath string) error

}

type FileRepository interface {
	// InsertOrUpdate(query string, values []interface{}) error
	CopyFrom(columns []string, values [][]interface{},tableName string,channel chan map[string]interface{},wg *sync.WaitGroup)
}