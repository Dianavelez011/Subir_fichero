package main

import (
	"log"
	"os"

	// "github.com/gin/cmd/api/handlers/file"
	"github.com/gin/cmd/api/routes"
	"github.com/gin/config"
	// "github.com/gin/config"
	"github.com/gin/internal/repositories/postgresql"
	"github.com/gin/internal/repositories/postgresql/connection"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil{
		log.Fatal("error loading .env file",err.Error())
	}

	connectPool,err := connection.Connect(os.Getenv("DATABASE_URL"))

	if err != nil{
		log.Fatal(err.Error())
	}

	repo := postgresql.Repository{
		Connection: connectPool,
	}

	handler := config.ConfigAllServices(repo)

	routes := routes.Router{
		Handler: handler,
	}

	routes.Routes();


	


}