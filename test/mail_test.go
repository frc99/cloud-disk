package test

import (
	"cloud-disk/core/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <fangrc0721@163.com>"
	e.To = []string{"fangrc0721@126.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "您的验证码为"
	//e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>验证码</h1>")
	err := e.SendWithTLS("smtp.163.com:587", smtp.PlainAuth("", "fangrc0721@163.com", define.MailPassword, "smtp.163.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
