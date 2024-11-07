package txt

import (
	"github.com/gin/internal/ports/file/txt"
	"golang.org/x/text/encoding"
)

type Service struct {
	Repo txt.FileRepository
	Decoder *encoding.Decoder
	FileLocation string
	Character string
	DefaultValue string
	TabRow int
	FilterWords []string
	Columns []string
	TableName string
	MainFileName string
}