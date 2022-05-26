package middleware

import (
	"bluebell/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
   return func(ctx *gin.Context) {
   	authHeader := ctx.Request.Header.Get("Authorization")
   	if authHeader == "" {
   		utils.ResponseError(ctx, utils.CodeUserTokenNotExist)
   		ctx.Abort()//结束后续操作
   		return
   	}

   	//按空格拆分
   	parts := strings.SplitN(authHeader, " ", 2)
   	if !(len(parts) == 2 && parts[0] == "Bearer")  {
   		utils.ResponseError(ctx, utils.CodeUserTokenError)
   		ctx.Abort()
   		return
   	}

   	//解析token包含的信息
   	claims ,err := utils.ParseToken(parts[1])
   	if err != nil {
   		utils.ResponseError(ctx, utils.CodeUserTokenError)
   		ctx.Abort()
   		return
   	}

   	// 将当前请求的claims.UserID信息保存到请求的上下文c上
   	ctx.Set("userid", claims.UserID)
   	ctx.Next() // 后续的处理函数可以用过ctx.Get("userid")来获取当前请求的用户信息
   }
}

