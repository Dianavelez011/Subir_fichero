package routes

import (
	"fmt"
	"github.com/gin/cmd/api/handlers/file"
)


type Router struct{
	Handler file.Handler
}

func (r Router)Routes() {
	ginEngine := r.config();
	routes := ginEngine.Group("/api")

	// routes.POST("/", createEvent)               // Crear un nuevo registro
    // routes.GET("/:cedula", getEventByCedula)
    routes.POST("/upload", r.Handler.Create)
    routes.POST("/delete", r.Handler.Delete)

	// Ejecutar el servidor
	if err := ginEngine.Run(":8081"); err != nil {
		fmt.Printf("Error al iniciar el servidor: %s", err.Error())
	}

}