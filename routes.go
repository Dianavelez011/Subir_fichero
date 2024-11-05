package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin/files"
)
func registerRoutes(e *gin.Engine) {
    victimas := e.Group("/api/v1/victimas")
    {
        victimas.POST("/", createEvent)               // Crear un nuevo registro
        victimas.GET("/:cedula", getEventByCedula)
        victimas.POST("/upload", func(ctx *gin.Context) {
            file, err := ctx.FormFile("file")
            if err != nil {
                ctx.JSON(http.StatusBadRequest, gin.H{"error": "file no found"})
                return
            }
            //Abrir archivo
            openedFile, err := file.Open()
    
            if err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unabled to open file"})
                return
            }
            //Cerrar archivo para ahorrar recursos
            defer openedFile.Close()
    
            buffer := make([]byte, 512)
    
            if _, err := openedFile.Read(buffer); err != nil && err != io.EOF {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read file"})
                return
            }
    
            //Obtener el tipo de archiovo
            mimetype := http.DetectContentType(buffer)
            fmt.Println("mimetype", mimetype)
            //Almacenar en un map los tipo de archivos permitidos
            allowedMimeTypes := map[string]bool{
                "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": true, // Excel .xlsx
                "application/vnd.ms-excel":  true, // Excel .xls
                "text/plain; charset=utf-8": true, //Txt
            }
    
            //Fallar si el tipo de archivo no esta permitido
            if !allowedMimeTypes[mimetype] {
                ctx.JSON(http.StatusUnsupportedMediaType, gin.H{"error": mimetype})
                return
            }
    
            if err := ctx.SaveUploadedFile(file, "./uploads/"+file.Filename); err != nil {
                ctx.JSON(http.StatusBadRequest, gin.H{"error": "file dont save"})
                return
            }
    
            if mimetype == "text/plain; charset=utf-8" {
                files.GetCsvData(file.Filename)
            } else {
                files.GetExcelData(file.Filename)
            }
    
            ctx.JSON(http.StatusOK, gin.H{"message": "file upload success!"})
        })
        
    }
}
