package controllers

import (
	"bluebell/models"
	"bluebell/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 获取参数并校验
	param := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("绑定参数失败", zap.Error(err))
		c.JSON(200, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	// 用户注册
	if err := services.SignUp(param); err != nil {
		zap.L().Error("用户注册失败", zap.Error(err))
		c.JSON(200, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "success"})
}
