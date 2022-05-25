package models

import (
	"bluebell/config"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func InitRedis(cfg *config.RedisConf) (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     cfg.Add,
		Password: cfg.Password,
		DB:       cfg.DBNum,
	})
	_, err = RDB.Ping(RDB.Context()).Result()
	return
}

func CloseRedis() {
	RDB.Close()
}