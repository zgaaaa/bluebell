package main

import (
	"bluebell/config"
	"bluebell/models"
	"bluebell/routes"
)

func main() {
	// 初始化数据库
	err := models.InitMysql(config.Conf.MysqlConf)
	if err != nil {
		panic(err)
	}
	defer models.CloseMysql()
	// 初始化缓存
	err = models.InitRedis(config.Conf.RedisConf)
	if err != nil {
		panic(err)
	}
	defer models.CloseRedis()
	// 初始化服务器
	routes.SetUp().Run(":8080")
}
