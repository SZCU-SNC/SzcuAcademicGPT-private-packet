package Util

import "regexp"

func CheckContactInfo(input string) string {
	phoneRegex := `^1[3456789]\d{9}$`
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if matched, _ := regexp.MatchString(phoneRegex, input); matched {
		return "phone"
	} else if matched, _ := regexp.MatchString(emailRegex, input); matched {
		return "email"
	} else {
		//todo 抛出错误
		return "未知格式"
	}
}
