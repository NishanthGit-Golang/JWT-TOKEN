package router

import (
	"jwt/controller"
	middleware "jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(incomeRoutes *gin.Engine) {
	incomeRoutes.Use(middleware.Authendicate())
	incomeRoutes.GET("/users", controller.GetUsers())
	incomeRoutes.GET("/users", controller.GetUserByID())
}
