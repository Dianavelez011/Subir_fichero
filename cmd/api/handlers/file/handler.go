package file

import (
	"github.com/gin/internal/ports/file"
	"github.com/gin/internal/ports/file/excel"
	"github.com/gin/internal/ports/file/txt"
)

type Handler struct {
	FileService file.FileService
	ExcelService excel.FileService
	TxtService txt.FileService
}