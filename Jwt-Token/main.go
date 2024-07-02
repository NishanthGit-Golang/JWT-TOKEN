package main

import (
	"jwt/router"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	rout := gin.New()
	rout.Use(gin.Logger())

	router.AuthRouters(rout)
	router.UserRouter(rout)

	rout.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": "Access Given For API-1",
		})
	})
	rout.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": "Access Given For API-2",
		})
	})

	rout.Run(":" + port)

}
