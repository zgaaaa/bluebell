package models

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

const (
	KeyPostPrefix = "bluebell:post:"
	KeyPostTime   = "time"
	KeyPostScore  = "score"
	KeyPostVote   = "vote:"

	score = 432
)

func GetPostKey(key string) string {
	return KeyPostPrefix + key
}

// CreatePostVote	创建投票
func CreatePostVote(postid int64) error {
	ctx := context.Background()
	pipeline := RDB.TxPipeline()
	pipeline.ZAdd(ctx, GetPostKey(KeyPostTime), &redis.Z{
		Member: postid,
		Score:  float64(time.Now().Unix()),
	})
	pipeline.ZAdd(ctx, GetPostKey(KeyPostScore), &redis.Z{
		Member: postid,
		Score:  float64(time.Now().Unix()),
	})
	_, err := pipeline.Exec(ctx)
	return err
}

// CreateVote 创建投票
func CreateVote(postid int64, userid int64, direction int8) error {
	ctx := context.Background()
	return RDB.ZAdd(ctx, GetPostKey(KeyPostVote+cast.ToString(postid)), &redis.Z{
		Score:  float64(direction),
		Member: userid,
	}).Err()
}

// GetVote 获取投票数
func GetVote(key string, member int64) float64 {
	ctx := context.Background()
	return RDB.ZScore(ctx, key, cast.ToString(member)).Val()
}

// UpdateVote 更新投票数
func UpdateVote(key string, postId int64, score float64) error {
	ctx := context.Background()
	return RDB.ZIncrBy(ctx, key, score, cast.ToString(postId)).Err()
}

// DelVote 删除投票
func DelVote(key string, member int64) error {
	ctx := context.Background()
	return RDB.ZRem(ctx, key, cast.ToString(member)).Err()
}
