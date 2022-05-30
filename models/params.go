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

// 投票参数
type ParamVote struct {
	UserId int64 `json:"user_id" db:"user_id"`
	PostId int64 `json:"post_id,string" binding:"required"`
	Direction int8 `json:"direction,omitempty" binding:"oneof=1 0 -1"` // 赞成票（1）、反对票（-1）、取消投票（0）
}

// 帖子列表参数
type ParamPostList struct {
	PageNum int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
	Order string `json:"order" form:"order"`
}

type ResPostList struct {
	Post
	Vote int64 `json:"votes"`
	Score float64 `json:"score"`	
}