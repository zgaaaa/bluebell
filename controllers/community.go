package controllers

import (
	"bluebell/services"
	"bluebell/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// CommunityListHandler 获取社区列表
func CommunityListHandler(c *gin.Context) {

	list, err := services.CommunityList()
	if err != nil {
		zap.L().Error("获取列表失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	// 返回结果
	utils.ResponseSuccess(c, list)
}

// CommunityDetailHandler 获取社区详情
func CommunityDetailHandler(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	data, err := services.CommunityDetail(id)
	if err != nil {
		zap.L().Error("获取详情失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	// 返回结果
	utils.ResponseSuccess(c, data)
}