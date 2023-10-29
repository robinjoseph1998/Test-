package routes

import (
	"Test/Handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, handler *Handlers.TestHandlers) *gin.Engine {

	r.POST("/addName", handler.Sample)
	r.GET("/showName", handler.ShowName)
	return r
}
