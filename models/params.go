package models

import "time"

// 用户注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required,min=4,max=16"`
	Password   string `json:"password" binding:"required,min=6,max=16"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

// 用户登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required,min=4,max=16"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

// 获取社区列表参数
type ParamCommunityList struct {
	Id   int    `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

// 社区详情参数
type ParamCommunityDetail struct {
	Id          int    `json:"id" db:"community_id"`
	Name        string `json:"name" db:"community_name"`
	Introduction string `json:"introduction" db:"introduction"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}