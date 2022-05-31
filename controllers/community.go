package controllers

import (
	"bluebell/models"
	"bluebell/services"
	"bluebell/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// CommunityCreateHandler 创建社区
func CommunityCreateHandler(c *gin.Context) {
	// 获取删除及参数校验
	community := &models.ParamCommunityCreate{}
	if err := c.ShouldBindJSON(community); err != nil {
		zap.L().Error("参数校验失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeParamError)
		return
	}
	// 创建新社区
	if err := services.CreateCommunity(community, c); err != nil {
		zap.L().Error("创建社区失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, nil)
	// 返回相应

}

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