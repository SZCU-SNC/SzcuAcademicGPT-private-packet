// Package EmailUtil
/*
配置文件：
email:
	host: smtp.qq.com
	port: 587
	auth:
		addr: your_email_account
		password: your_authorization_code
*/
package EmailUtil

import (
	"fmt"
	"github.com/SZCU-SNC/SzcuAcademicGPT-private-packet/Utils/ConfigUtil"
	"gopkg.in/gomail.v2"
	"os"
)

func SendEmail(to []string, subject string, context string, filePath string) error {

	host := ConfigUtil.GetConfigData()["email"].(map[interface{}]interface{})["host"].(string)
	port := ConfigUtil.GetConfigData()["email"].(map[interface{}]interface{})["port"].(int)

	auth := ConfigUtil.GetConfigData()["email"].(map[interface{}]interface{})["auth"].(map[interface{}]interface{})
	addr := auth["addr"].(string)
	password := auth["password"].(string)

	e := gomail.NewMessage()
	//设置发送方的邮箱
	e.SetHeader("From", addr)
	e.SetHeader("To", to...)

	e.SetHeader("Subject", subject)

	e.SetBody("text/html", context)

	if filePath != "" {
		_, err := os.Stat(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				return err // 文件不存在
			}
		}
		//副件
		e.Attach(filePath)
	}

	d := gomail.NewDialer(host, port, addr, password)

	if err := d.DialAndSend(e); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}

	return nil
}
