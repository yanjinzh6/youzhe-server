package email

import (
	// "encoding/base64"
	"github.com/yanjinzh6/youzhe-server/tools"
	"net/smtp"
	"strings"
)

func SendEmail(username, password, host, to, subject, body, mailType string) error {
	hs := strings.Split(host, ":")
	auth := smtp.PlainAuth("", username, password, hs[0])
	tools.Println(auth)
	contentType := "Content-Type: text/" + mailType + "; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: " + username + "<" + username + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ",")
	tools.Println(sendTo)
	err := smtp.SendMail(host, auth, username, sendTo, msg)
	tools.Println(err)
	return err
}
