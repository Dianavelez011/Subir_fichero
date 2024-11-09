package txt

import (
	"context"
	"mime/multipart"
	"sync"
)

type FileService interface {
	Create(ctx context.Context, file *multipart.FileHeader, sizeMainFile int,channel chan map[string]interface{},wg *sync.WaitGroup)
	DeleteFolder() error

}

type FileRepository interface {
	// InsertOrUpdate(query string, values []interface{}) error
	CopyFrom(columns []string, values [][]interface{},tableName string,channel chan map[string]interface{},wg *sync.WaitGroup)
}