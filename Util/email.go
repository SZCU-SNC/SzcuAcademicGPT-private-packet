package Util

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func SendEmail(to []string, subject string, context string, filePath string) {

	server := GetConfigData()["email"].(map[interface{}]interface{})["server"].(string)
	auth := GetConfigData()["email"].(map[interface{}]interface{})["auth"].(map[interface{}]interface{})
	username := auth["username"].(string)
	password := auth["password"].(string)
	host := auth["host"].(string)

	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = username

	e.To = to
	e.Subject = subject
	e.Text = []byte(context)

	//副件
	_, err := e.AttachFile(filePath)
	if err != nil {
		return
	}

	// 使用配置值进行发送邮件操作
	err = e.Send(server, smtp.PlainAuth("", username, password, host))

	if err != nil {
		log.Fatal(err)
	}
}
