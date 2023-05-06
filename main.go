package main

import (
	"github.com/andyanunciacion/golang-basic-crud/initializers"
	"github.com/andyanunciacion/golang-basic-crud/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	routes.SetPostRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}