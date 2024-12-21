package email

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"sync"
)

// 由于email库调用send函数的时候，需要将地址，认证信息参数传入进去，这里进行封装一下
type EmailSender struct {
	rw       sync.RWMutex
	Email    *email.Email
	Host     string
	Port     string
	Username string
	Password string
	TLS      bool
}

func NewEmailSender(sendEmail string, sendSubject string, Username string, Password string, host string, port string, tls bool) *EmailSender {
	e := email.NewEmail()
	e.From = sendEmail
	e.Subject = sendSubject

	return &EmailSender{
		Email:    e,
		Host:     host,
		Port:     port,
		Username: Username,
		Password: Password,
		TLS:      tls,
	}
}

func (e *EmailSender) SendSampleCode(code string, acceptEmail string) error {
	e.rw.Lock()
	defer e.rw.Unlock()
	e.Email.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	e.Email.To = []string{acceptEmail}
	
	addr := e.Host + ":" + e.Port
	fmt.Println("addr", addr)
	err := e.Email.SendWithTLS(addr, smtp.PlainAuth("", e.Username, e.Password, e.Host),
		&tls.Config{InsecureSkipVerify: e.TLS, ServerName: e.Host})
	return err
}
