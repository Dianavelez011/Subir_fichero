package excel

import (
	"github.com/gin/internal/ports/file/excel"
)

type Handler struct {
	FileService excel.FileService
}