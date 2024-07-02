package router

import (
	"jwt/controller"

	"github.com/gin-gonic/gin"
)

func AuthRouters(incoming *gin.Engine) {
	incoming.POST("users/singup", controller.SignUp())
	incoming.POST("users/login", controller.Login())

}
