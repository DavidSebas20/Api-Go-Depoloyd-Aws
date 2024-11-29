package main

import "github.com/gin-gonic/gin"

func main() {

	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Api Go from Docker deployd on Aws",
		})
	})

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run()
}
