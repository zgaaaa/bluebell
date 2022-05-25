package routes

import (
	"bluebell/middleware"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger(), middleware.GinRecovery(true))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
