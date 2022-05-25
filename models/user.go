package models

import (
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) (bool, error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := DB.Get(&count, sqlStr, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

// InsertUser 插入用户
func InsertUser(user *User) error {
	// 对密码加密
	user.Password = encryptPassword(user.Password)
	// 插入数据库
	sqlStr := "insert into user(user_id, username, password, email) values(:user_id, username, password, email)"
	_, err := DB.NamedExec(sqlStr, user)
	return err
}

func encryptPassword(rawpassword string) string {
	ps := md5.New()
	ps.Write([]byte(rawpassword))
	return hex.EncodeToString(ps.Sum([]byte(rawpassword)))
}