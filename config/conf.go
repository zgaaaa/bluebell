package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConf)

type AppConf struct {
	*ServerConf    `mapstructure:"server"`
	*MysqlConf     `mapstructure:"mysql"`
	*RedisConf     `mapstructure:"redis"`
	*JwtConf       `mapstructure:"jwt"`
	*LogConf       `mapstructure:"log"`
	*SnowflakeConf `mapstructure:"snowflake"`
}

type ServerConf struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type MysqlConf struct {
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	DBName       string `mapstructure:"dbname"`
	MaxIdleConns int    `mapstructure:"maxidleconns"`
	MaxOpenConns int    `mapstructure:"maxopenconns"`
}

type RedisConf struct {
	Add      string `mapstructure:"add"`
	DBNum    int    `mapstructure:"dbnum"`
	Password string `mapstructure:"password"`
}

type LogConf struct {
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxsize"`
	MaxAge     int    `mapstructure:"maxage"`
	MaxBackups int    `mapstructure:"maxbackups"`
	Level      string `mapstructure:"level"`
}

type JwtConf struct {
	Key string `mapstructure:"key"`
}

type SnowflakeConf struct {
	MachineId int64  `mapstructure:"machineid"`
	StartTime string `mapstructure:"starttime"`
}

func init() {
	// 获取当前文件所在目录
	file := "config/conf.yaml"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	viper.SetConfigFile(file) // 配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 将配置文件内容解析到结构体
	if err := viper.Unmarshal(Conf); err != nil {
		panic(err)
	}
	viper.WatchConfig() // 监控配置文件变化
	// 当配置文件发生变化时，重新读取配置文件
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("配置文件解析错误: %v\n", err)
		}
	})

}
