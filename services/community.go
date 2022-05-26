package services

import "bluebell/models"

// CommunityList 获取社区列表
func CommunityList() ([]*models.ParamCommunityList, error) {
	return models.CommunityList()
}

// CommunityDetail 获取社区详情
func CommunityDetail(id int) (*models.ParamCommunityDetail, error) {
	return models.CommunityDetail(id)
}

