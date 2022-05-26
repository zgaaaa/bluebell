package controllers

import (
	"bluebell/models"
	"bluebell/services"
	"bluebell/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// SignUpHandler 用户注册
func SignUpHandler(c *gin.Context) {
	// 获取参数并校验
	param := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("绑定参数失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeParamError)
		return
	}
	// 用户注册
	if err := services.SignUp(param); err != nil {
		zap.L().Error("用户注册失败", zap.Error(err))
		if err == services.ErrUserExist {
			utils.ResponseError(c, utils.CodeUserExist)
		} else {
			utils.ResponseError(c, utils.CodeServerBusy)
		}
		return
	}
	utils.ResponseSuccess(c, nil)
}

// LoginHandler 用户登录
func LoginHandler(c *gin.Context) {
	// 获取参数并校验
	param := new(models.ParamLogin)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("绑定参数失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeParamError)
		return
	}
	// 用户登录
	token, err := services.Login(param)
	if err != nil {
		zap.L().Error("用户登录失败", zap.Error(err))
		if err == services.ErrUserNotExist {
			utils.ResponseError(c, utils.CodeUserNotExist)
		} else if err == services.ErrPassword {
			utils.ResponseError(c, utils.CodeUserPasswordError)
		} else {
			utils.ResponseError(c, utils.CodeServerBusy)
		}
		return
	}
	// 登录成功
	utils.ResponseSuccess(c, token)
}
