package services

import (
	"bluebell/models"
	"bluebell/utils"
	"errors"
)


// SignUp 用户注册
func SignUp(p *models.ParamSignUp) error {
	// 检查用户是否存在
	if exist, err := models.CheckUserExist(p.Username); err != nil {
		return err
	} else if exist {
		return errors.New("用户名已存在")
	}
	// 创建用户实例
	user := &models.User{
		UserID:   utils.GetID(),
		Username: p.Username,
		Password: p.Password,
	}
	// 插入数据库
	return models.InsertUser(user)
}
