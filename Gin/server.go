package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "EVERYTHING GOOOD! Goland VS code extension needs serious upkeep"})
	})

	server.Run(":8080")

}
