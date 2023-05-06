package controllers

import (
	"github.com/andyanunciacion/golang-basic-crud/initializers"
	"github.com/andyanunciacion/golang-basic-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Get data from req body
	var body struct{
		Title string
		Body string
	}
	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	// Check if the create resulted in error
	if result.Error != nil{
		c.Status(400)
		return
	}

	// Return result(post you've created)
	c.JSON(200,gin.H{
		"post": post,
	})
}


func PostsIndex(c *gin.Context) {

	// Fetch Posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return fetched posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get params id
	id := c.Param("id")

	// Fetch Post By Id
	var post models.Post
	initializers.DB.First(&post, id)

	// Return fetched post
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostUpdate (c *gin.Context) {
	// Get params id
	id := c.Param("id")

	// Get data from req body
	var body struct{
		Title string
		Body string
	}
	c.Bind(&body)

	// Fetch Post By Id
	var post models.Post
	initializers.DB.First(&post, id)

	// Update Post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body:body.Body})

	// Return updated post
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context){
	// Get params id
	id := c.Param("id")

	// Delete post
	initializers.DB.Delete(&models.Post{}, id)

	// Return message
	c.JSON(200, gin.H{
		"message": "Post has been deleted",
	})
}






