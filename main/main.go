package main

import (
	"awesomeProject/main/domain"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(CORSMiddleware())
	api := r.Group("/api")
	{
		api.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{"message": "pong"})
		})
		api.GET("/albums", domain.GetAlbum)
		api.POST("/albums", domain.PostAlbums)
	}

	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
