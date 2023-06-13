package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"proyecto.com/Infraestructure/Database"
	"proyecto.com/Router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error al cargar las variables de entorno", err)
		os.Exit(1)
	}
	Database.Connection()

	router := gin.Default()
	Router.PlantRoutes(router)

	router.Run(":5000")
}
