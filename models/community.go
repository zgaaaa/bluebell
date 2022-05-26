package models

import (
	"database/sql"

	"go.uber.org/zap"
)

// CommunityList 获取社区列表
func CommunityList() ([]*ParamCommunityList, error) {
	communitys := make([]*ParamCommunityList, 0)
	sqlStr := "select community_id, community_name from community"
	err := DB.Select(&communitys, sqlStr)
	if err != sql.ErrNoRows {
		zap.L().Warn("没有查询到数据", zap.Error(err))
		err = nil
	}
	return communitys, err
}

// CommunityDetail 获取社区详情
func CommunityDetail(id int) (*ParamCommunityDetail, error) {
	community := new(ParamCommunityDetail)
	sqlStr := "select community_id, community_name, introduction, create_time from community where community_id = ?"
	err := DB.Get(community, sqlStr, id)
	if err != nil {
		zap.L().Warn("没有查询到数据", zap.Error(err))
		err = nil
	}
	return community, err
}
