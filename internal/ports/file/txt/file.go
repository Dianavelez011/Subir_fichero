package txt

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type FileService interface {
	Create(ctx *gin.Context,file *multipart.FileHeader,sizeMainFile int )error
}

type FileRepository interface {
	// InsertOrUpdate(query string, values []interface{}) error
	CopyFrom(columns []string, values [][]interface{},tableName string) error
}