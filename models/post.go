package models

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	OrderByTime  = "time"
	OrderByScore = "score"
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

func GetPostIdsByOrder(param *ParamPostList) ([]string, error) {
	// 根据param.Order参数选择排序方式
	key := GetPostKey(KeyPostTime)
	if param.Order == OrderByScore {
		key = GetPostKey(KeyPostScore)
	}
	// 查询范围
	start := (param.PageNum - 1) * param.PageSize
	end := start + param.PageSize - 1
	// 按分数从大到小查询
	return RDB.ZRevRange(context.Background(), key, int64(start), int64(end)).Result()
}

func GetPostsByIds(ids []string) ([]*ResPostList, error) {
	sqlStr := `SELECT post_id, title, content, author_id, community_id, status, create_time
	FROM post 
	WHERE post_id IN (?)
	ORDER BY FIND_IN_SET(post_id, ?)`
	posts := make([]*ResPostList, 0, len(ids))
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = DB.Rebind(query)
	if err := DB.Select(&posts, query, args...); err != nil {
		return nil, err
	}

	votes,scores, err := GetPostsScore(ids)
	if err != nil {
		return nil, err
	}
	for i, post := range posts {
		post.Vote = votes[i]
		post.Score = scores[i]
	}
	return posts, nil
}
