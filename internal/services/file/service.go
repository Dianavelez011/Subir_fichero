package file

import (
	"github.com/gin/internal/ports/file/excel"
	"github.com/gin/internal/ports/file/txt"
)

type Service struct {
	Repo excel.FileRepository
	ExcelService excel.FileService
	TxtService  txt.FileService
}