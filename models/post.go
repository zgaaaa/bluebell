package models

import (
	"database/sql"
	"time"
)

type Post struct {
	Id          int64     `json:"id" db:"post_id"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	AuthorId    int64     `json:"author_id" db:"author_id"`
	CommunityId int       `json:"community_id" db:"community_id" binding:"required"`
	Status      int       `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// CreatePost 创建帖子
func CreatePost(post *Post) error {
	sqlStr := `INSERT INTO post 
	(post_id, title, content, author_id, community_id) 
	VALUES (?, ?, ?, ?, ?)`
	_, err := DB.Exec(sqlStr, post.Id, post.Title, post.Content, post.AuthorId, post.CommunityId)
	return err
}

// GetPostDetail 获取帖子详情
func GetPostDetail(id int64) (*Post, error) {
	sqlStr := `SELECT post_id, title, content, author_id, community_id, status, create_time 
	FROM post WHERE post_id = ?`
	post := &Post{}
	err := DB.Get(post, sqlStr, id)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return post, nil
}
