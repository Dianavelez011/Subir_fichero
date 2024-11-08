package txt

import (
	"context"
	"mime/multipart"
	"sync"
)

type FileService interface {
	Create(ctx context.Context, file *multipart.FileHeader, sizeMainFile int,channel chan error,wg *sync.WaitGroup)
}

type FileRepository interface {
	// InsertOrUpdate(query string, values []interface{}) error
	CopyFrom(columns []string, values [][]interface{},tableName string) error
}