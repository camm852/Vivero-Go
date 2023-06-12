package main

import (
	"github.com/gin-gonic/gin"
	"proyecto.com/Router"
)

func main() {
	router := gin.Default()
	Router.PlantRoutes(router)

	router.Run(":5000")
}
