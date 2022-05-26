package services

import (
	"bluebell/models"
	"bluebell/utils"
	"errors"
)

var (
	ErrUserNotExist = errors.New("用户不存在")
	ErrUserExist   = errors.New("用户已存在")
	ErrPassword	= errors.New("用户名或密码错误")
)


// SignUp 用户注册
func SignUp(p *models.ParamSignUp) error {
	// 检查用户是否存在
	if exist, err := models.CheckUserExist(p.Username); err != nil {
		return err
	} else if exist {
		return ErrUserExist
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

// Login 用户登录
func Login(p *models.ParamLogin) error {
	// 检查用户是否存在
	if exist, err := models.CheckUserExist(p.Username); err != nil {
		return err
	} else if !exist {
		return ErrUserNotExist
	}
	// 检查密码
	if pwd, err := models.GetUserPassword(p.Username); err != nil {
		return err
	} else if pwd != models.EncryptPassword(p.Password) {
		return ErrPassword
	}
	return nil
}
