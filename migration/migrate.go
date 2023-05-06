package main

import (
	"github.com/andyanunciacion/golang-basic-crud/initializers"
	"github.com/andyanunciacion/golang-basic-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}