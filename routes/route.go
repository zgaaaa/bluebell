package routes

import (
	"bluebell/controllers"
	"bluebell/middleware"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger(), middleware.GinRecovery(true))

	noAuth := r.Group("/api/v1") // 无需鉴权的接口
	{
		noAuth.POST("/signup", controllers.SignUpHandler)
		noAuth.POST("/login", controllers.LoginHandler)
	}
	auth := r.Group("/api/v1") // 需要鉴权的接口
	{
		auth.Use(middleware.JWTAuth())
		auth.GET("/community/list", controllers.CommunityListHandler)
		auth.GET("/community/:id", controllers.CommunityDetailHandler)

		auth.POST("/post", controllers.PostCreateHandler)
		auth.GET("/post/:id", controllers.PostDetailHandler)
		auth.GET("/post/list", controllers.PostListHandler)

		auth.POST("/vote", controllers.VoteHandler)
	}
	return r
}
