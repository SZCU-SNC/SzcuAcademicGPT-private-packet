package RedisKey

import (
	"strconv"
)

// RegisterVerificationCode 注册用邮箱、手机验证码
func RegisterVerificationCode(user string) string {
	return "VerificationCode:Register:" + user
}

// UserToken 用户token
func UserToken(userId string) string {
	return "UserToken:" + userId
}

// CaptchaCode 图形验证码
func CaptchaCode(user string) string {
	return "CaptchaCode:" + user
}

func UserRole(userId int) string {
	return "UserRole:" + strconv.Itoa(userId)
}
