package models

// 用户注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required,min=4,max=16"`
	Password   string `json:"password" binding:"required,min=6,max=16"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

// 用户登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required,min=4,max=16"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}