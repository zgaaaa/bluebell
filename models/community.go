package models

import (
	"database/sql"

	"go.uber.org/zap"
)

const (
	KeyCommunity = "bluebell:community:"
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

// CreateCommunity 创建社区
func CreateCommunity(community *ParamCommunityCreate) error {
	sqlStr := "insert into community (community_id, community_name, introduction) values (?, ?, ?, ?)"
	_, err := DB.Exec(sqlStr, community.Id, community.Name, community.Introduction)
	return err
}