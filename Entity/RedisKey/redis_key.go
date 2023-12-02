package RedisKey

import (
	"strconv"
)

// RegisterVerificationCode 注册用邮箱、手机验证码
func RegisterVerificationCode(user string) string {
	return "VerificationCode:Register:" + user
}

// CaptchaCode 图形验证码
func CaptchaCode(userId int) string {
	return "CaptchaCode:" + strconv.Itoa(userId)
}
