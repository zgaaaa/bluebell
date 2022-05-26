package utils

import "github.com/gin-gonic/gin"

type ResCode int

const (
	CodeSuccess           ResCode = 1000 + iota
	CodeParamError                // 参数错误
	CodeUserNotExist              // 用户不存在
	CodeUserExist                 // 用户已存在
	CodeUserPasswordError         // 用户密码错误
	CodeUserTokenError            // 用户token错误
	CodeUserTokenExpire           // 用户token过期
	CodeUserTokenNotExist         // 用户token不存在
	CodeUserTokenNotMatch         // 用户token不匹配
	CodeUserTokenNotValid         // 用户token不合法
	CodeServerBusy                // 服务器繁忙
	
)

var CodeMsg = map[ResCode]string{
	CodeSuccess:           "success",
	CodeParamError:        "请求参数错误",
	CodeUserNotExist:      "用户不存在",
	CodeUserExist:         "用户已存在",
	CodeUserPasswordError: "用户或密码错误",
	CodeUserTokenError:    "用户token错误",
	CodeUserTokenExpire:   "用户token过期",
	CodeUserTokenNotExist: "用户token不存在",
	CodeUserTokenNotMatch: "用户token不匹配",
	CodeUserTokenNotValid: "用户token不合法",
	CodeServerBusy:        "服务器繁忙",

}

type Response struct {
	Code    ResCode `json:"code"`
	Message string  `json:"message"`
	Data    any     `json:"data"`
}

func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(200, Response{
		Code:    CodeSuccess,
		Message: CodeMsg[CodeSuccess],
		Data:    data,
	})
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(200, Response{
		Code:    code,
		Message: CodeMsg[code],
		Data:    nil,
	})
}

func ResponseErrorMsg(c *gin.Context, code ResCode, msg string) {
	c.JSON(200, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}