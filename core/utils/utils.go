package utils

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenSaltPassword(salt, password string) string {
	s1 := sha256.New()
	s1.Write([]byte(password))
	str1 := fmt.Sprintf("%x", s1.Sum(nil))

	s2 := sha256.New()
	s2.Write([]byte(str1 + salt))
	str2 := fmt.Sprintf("%x", s2.Sum(nil))
	return str2
}
func GenerateToken(id int, identity, name string) (string, error) {
	//id
	// identity
	// name
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "验证码 <fangrc0721@163.com>"
	e.To = []string{mail}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "您的验证码为" + code
	//e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte(code)
	err := e.SendWithTLS("smtp.163.com:587", smtp.PlainAuth("", "fangrc0721@163.com", define.MailPassword, "smtp.163.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func GetUUID() string {
	return uuid.NewV4().String()
}
