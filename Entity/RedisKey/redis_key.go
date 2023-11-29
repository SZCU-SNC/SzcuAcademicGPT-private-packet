package RedisKey

import (
	"strconv"
)

// VerificationCode 邮箱、手机验证码
func VerificationCode(userId int) string {
	return "VerificationCode:" + strconv.Itoa(userId)
}

// CaptchaCode 图形验证码
func CaptchaCode(userId int) string {
	return "CaptchaCode:" + strconv.Itoa(userId)
}
