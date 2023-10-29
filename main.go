package main

import (
	routes "Test/routes"

	"Test/di"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	handlers := di.InitializeHandlers()
	routes.SetupRoutes(r, handlers)
	r.Run(":8000")

}
