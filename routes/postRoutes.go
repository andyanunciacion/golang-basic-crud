package routes

import (
	"github.com/andyanunciacion/golang-basic-crud/controllers"
	"github.com/gin-gonic/gin"
)

func SetPostRoutes(router *gin.Engine){
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.POST("/posts", controllers.PostsCreate)
	router.PUT("/posts/:id", controllers.PostUpdate)
	router.DELETE("/posts/:id", controllers.PostDelete)
}