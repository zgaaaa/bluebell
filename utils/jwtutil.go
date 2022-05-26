package utils

import (
	"bluebell/config"
	"bluebell/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type CustomClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims

}

var mysecret = []byte(config.Conf.JwtConf.Key)

// 创建 Token
func GenToken(user *models.User) (string, error) {
	claim := CustomClaims{
		user.UserID,
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 5)), //5分钟后过期
			Issuer:    config.Conf.Version,                     //签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(mysecret)
}

// 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	mc := new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenStr, mc, func(token *jwt.Token) (interface{}, error) {
		return mysecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验 token 
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// 如果token没有过期，延长token的过期时间
func RefreshToken(tokenStr string) (string, error) {
	mc, err := ParseToken(tokenStr)
	if err != nil {
		return "", err
	}
	mc.ExpiresAt = jwt.At(time.Now().Add(time.Minute * 10)) //10分钟后过期
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	return token.SignedString(mysecret)
}
