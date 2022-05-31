package services

import (
	"bluebell/models"
	"bluebell/utils"

	"github.com/gin-gonic/gin"
)

// CommunityList 获取社区列表
func CommunityList() ([]*models.ParamCommunityList, error) {
	return models.CommunityList()
}

// CommunityDetail 获取社区详情
func CommunityDetail(id int) (*models.ParamCommunityDetail, error) {
	return models.CommunityDetail(id)
}


// CreateCommunity 创建社区
func CreateCommunity(community *models.ParamCommunityCreate, c *gin.Context) error {
	community.Id = utils.GetID()
	return models.CreateCommunity(community)
}
