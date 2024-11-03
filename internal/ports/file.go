package ports

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type FileService interface {
	Create(ctx *gin.Context, file *multipart.FileHeader)
}