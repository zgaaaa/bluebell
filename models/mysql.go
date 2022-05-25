package models

import (
	"bluebell/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var DB *sqlx.DB

func InitMysql(cfg *config.MysqlConf) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("mysql connect error", zap.Error(err))
		return
	}
	DB.SetMaxIdleConns(cfg.MaxIdleConns)
	DB.SetMaxOpenConns(cfg.MaxOpenConns)
	return
}

func CloseMysql() {
	DB.Close()
}
