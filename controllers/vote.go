package controllers

import (
	"bluebell/models"
	"bluebell/services"
	"bluebell/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func VoteHandler(c *gin.Context) {
	p := new(models.ParamVote)
	if err := c.BindJSON(p); err != nil {
		zap.L().Error("bind json error", zap.Error(err))
		utils.ResponseError(c,utils.CodeParamError)
		return
	}
	if err := services.PostVote(c, p); err != nil {
		zap.L().Error("post vote error", zap.Error(err))
		utils.ResponseError(c,utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c,nil)
}