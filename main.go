package main

import (
	"github.com/gin-gonic/gin"
	"proyecto.com/Infraestructure/Database"
	"proyecto.com/Router"
)

func main() {

	Database.Connection()

	router := gin.Default()
	Router.PlantRoutes(router)

	router.Run(":5000")
}
