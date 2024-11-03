package file

import "github.com/gin/internal/ports"

type Handler struct {
	FileService ports.FileService
}