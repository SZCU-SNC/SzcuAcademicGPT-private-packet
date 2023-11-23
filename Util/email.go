package Util

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func (receiver Config) name(to []string, subject string, context string) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "dj <XXX@qq.com>"

	e.To = to
	e.Subject = subject
	e.Text = []byte(context)

	//副件
	_, err := e.AttachFile("./test.txt")
	if err != nil {
		return
	}
	//设置服务器相关的配置
	err = e.Send("smtp.qq.com:25", smtp.PlainAuth("", "你的邮箱账号", "这块是你的授权码", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
