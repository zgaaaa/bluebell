package services

import (
	"bluebell/models"
	"bluebell/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post, c *gin.Context) error {
	post.Id = utils.GetID()
	id, ok := c.Get("userid")
	if !ok {
		return errors.New("获取用户ID失败")
	}
	post.AuthorId = id.(int64)
	if err := models.CreatePost(post); err != nil {
		return err
	}
	// 创建投票
	if err := models.CreatePostVote(post.Id); err != nil {
		return err
	}
	return models.CreateCommunityPost(post.CommunityId, post.Id)
	
}

// GetPostDetail 获取帖子详情
func GetPostDetail(id int64) (*models.Post, error) {
	post, err := models.GetPostDetail(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetPostList 获取帖子列表
func GetPostList(param *models.ParamPostList) ([]*models.ResPostList, error) {
	ids, err := models.GetPostIdsByOrder(param)
	if err != nil {
		return nil, err
	}
	return models.GetPostsByIds(ids)

}


// GetPostListByCommunity 按社区获取帖子列表
func GetPostListByCommunity(param *models.ParamPostListByCommunity) ([]*models.ResPostList, error) {
	ids, err := models.GetPostIdsByCommunity(param)
	if err != nil {
		return nil, err
	}
	return models.GetPostsByIds(ids)
} 