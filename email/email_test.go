package email

import (
	"github.com/yanjinzh6/youzhe-server/tools"
	"testing"
)

func TestSendEmail(t *testing.T) {
	tools.Println(SendEmail("3316642293@qq.com", "1=2-30z/x.c,", "smtp.qq.com:25", "q32@foxmail.com", "test email", "go net/smtp qq mail testing", "plain"))
}
