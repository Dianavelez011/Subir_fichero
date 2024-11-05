package excel

import (
	"github.com/gin/internal/ports/file/excel"
)

type Service struct {
	Repo excel.FileRepository
	FileLocation string
	Columns []string
	TableName string
}