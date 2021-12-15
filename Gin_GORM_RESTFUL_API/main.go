package main

import (
	"api_intgr/controllers"
	"api_intgr/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	//})

	models.ConnectDatabase()
	server.GET("/books", controllers.FindBooks)

	server.POST("/books", controllers.CreateBook)

	server.GET("/books/:id", controllers.FindBook)
	server.PATCH("/books/:id", controllers.UpdateBook)
	server.DELETE("/books/:id", controllers.DeleteBook)
	err := server.Run()
	if err != nil {
		fmt.Println("Error with runnign server")
	}
}
