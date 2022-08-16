package routes

import (
	"github.com/gin-gonic/gin"
	controllers "gin/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/", controllers.IndexGetHandler())
	g.POST("/create", controllers.CreateItemHandler())
}