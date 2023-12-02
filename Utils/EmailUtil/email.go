package EmailUtil

import (
	"github.com/SZCU-SNC/SzcuAcademicGPT-private-packet/Utils/ConfigUtil"
	"github.com/jordan-wright/email"
	"net/smtp"
	"os"
)

func SendEmail(to []string, subject string, context string, filePath string) error {

	server := ConfigUtil.GetConfigData()["email"].(map[interface{}]interface{})["server"].(string)
	auth := ConfigUtil.GetConfigData()["email"].(map[interface{}]interface{})["auth"].(map[interface{}]interface{})
	username := auth["username"].(string)
	password := auth["password"].(string)
	host := auth["host"].(string)

	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = username

	e.To = to
	e.Subject = subject
	e.Text = []byte(context)

	if filePath != "" {
		_, err := os.Stat(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				return err // 文件不存在
			}
		}
		//副件
		_, err = e.AttachFile(filePath)
		if err != nil {
			return err
		}
	}

	// 使用配置值进行发送邮件操作
	err := e.Send(server, smtp.PlainAuth("", username, password, host))

	if err != nil {
		return err
	}

	return nil
}
