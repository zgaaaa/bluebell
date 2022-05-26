package routes

import (
	"bluebell/controllers"
	"bluebell/middleware"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger(), middleware.GinRecovery(true))
	r.POST("/signup", controllers.SignUpHandler)
	r.POST("/login", controllers.LoginHandler)
	return r
}
