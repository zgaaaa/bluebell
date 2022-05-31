package controllers

import (
	"bluebell/models"
	"bluebell/services"
	"bluebell/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// PostCreateHandler 创建帖子
func PostCreateHandler(c *gin.Context) {
	// 获取删除及参数校验
	post := &models.Post{}
	if err := c.ShouldBindJSON(post); err != nil {
		zap.L().Error("参数校验失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeParamError)
		return
	}
	// 创建新帖子
	if err := services.CreatePost(post, c); err != nil {
		zap.L().Error("创建帖子失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, nil)
	// 返回相应
}

// PostDetailHandler 获取帖子详情
func PostDetailHandler(c *gin.Context) {
	// 获取删除及参数校验
	id := cast.ToInt64(c.Param("id"))

	// 获取帖子详情
	post, err := services.GetPostDetail(id)
	if err != nil {
		zap.L().Error("获取帖子详情失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, post)
}

// PostListHandler 获取帖子列表
func PostListHandler(c *gin.Context) {
	// 获取删除及参数校验
	param := &models.ParamPostList{
		PageNum:  1,
		PageSize: 10,
		Order:    models.OrderByTime,
	}
	if err := c.ShouldBindQuery(param); err != nil {
		zap.L().Error("参数校验失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeParamError)
		return
	}
	// 获取帖子列表
	posts, err := services.GetPostList(param)
	if err != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, posts)
}

func PostListByCommunityHandler(c *gin.Context) {
	param := &models.ParamPostListByCommunity{
		CommunityId: cast.ToInt64(c.Param("id")),
		PageNum:     1,
		PageSize:    10,
		Order:       models.OrderByTime,
	}
	if err := c.ShouldBindQuery(param); err != nil {
		zap.L().Error("参数校验失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeParamError)
		return
	}
	// 获取帖子列表
	posts, err := services.GetPostListByCommunity(param)
	if err != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(err))
		utils.ResponseError(c, utils.CodeServerBusy)
		return
	}
	utils.ResponseSuccess(c, posts)
}
