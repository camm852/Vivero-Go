package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"proyecto.com/Application/Services"
	"proyecto.com/Infraestructure/Database"
	"proyecto.com/Router"
	"time"
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

	go decreaseNutrientsAndWater()

	router.Run(":5000")
}

func decreaseNutrientsAndWater() {
	for {
		time.Sleep(time.Second * 2)
		Services.DecreaseLife()
		fmt.Println("Decresio")
	}
}
