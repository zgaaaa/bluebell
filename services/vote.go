package services

import (
	"bluebell/models"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func PostVote(c *gin.Context, p *models.ParamVote) (err error) {
	userId,ok := c.Get("userid")
	if !ok {
		return errors.New("userid not found")
	}
	p.UserId = userId.(int64)
	// 投票的几种情况：
	// 1. vote=1：
	// 之前没有投过票，现在投赞成票     票差：1  +432
	// 之前投反对票，现在投赞成票       票差：2  +432*2
	// 2. vote=0：
	// 之前投反对票，现在取消投票       票差：1  +432
	// 之前投赞成票，现在取消投票       票差：-1  -432*2
	// 3. vote=-1：
	// 之前没有投过票，现在投反对票     票差：-1  -432
	// 之前投赞成票，现在投投反对票     票差：-2  -432*2
	keyName := models.GetPostKey(models.KeyPostVote+cast.ToString(p.PostId))
	oldDir := models.GetVote(keyName, p.UserId)
	sorce := (float64(p.Direction) - oldDir) * 432
	if err := models.UpdateVote(models.GetPostKey(models.KeyPostScore), p.PostId, sorce); err != nil {
		return err
	}
	if p.Direction == 0 {
		err = models.DelVote(keyName, p.UserId)
	} else {
		err = models.CreateVote(p.PostId,p.UserId,p.Direction)
	}
	return err
}