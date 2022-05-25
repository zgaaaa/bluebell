package services

import (
	"bluebell/models"
	"errors"
)

func SignUp(p *models.ParamSignUp) error {
	// 检查用户是否存在
	if exist, err := models.CheckUserExist(p.Username); err != nil {
		return err
	} else if exist {
		return errors.New("用户名已存在")
	}
	// 创建用户
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := models.InsertUser(user); err != nil {
		return err
	}
	return nil
}