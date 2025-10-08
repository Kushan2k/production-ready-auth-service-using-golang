package utils

import (
	"crypto/tls"

	"github/go_auth_api/internal/config"

	"gopkg.in/gomail.v2"
)


func  GetMailer() *gomail.Dialer {

	d := gomail.NewDialer(config.Envs.MAIL_HOST, config.Envs.MAIL_PORT, config.Envs.MAIL_USERNAME, config.Envs.MAIL_PASSWORD)
	
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return d
}

func GetMessage() *gomail.Message {
	return gomail.NewMessage()
}

func SendMail(body string, to string, subject string,res_chan_error chan<- bool) {
	mailer := GetMailer()
	msg := GetMessage()

	msg.SetBody("text/html", body)
	msg.SetHeader("From", config.Envs.MAIL_USERNAME)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)

	err:=mailer.DialAndSend(msg)
	if err != nil {
		res_chan_error<-true
	}else{
		res_chan_error<-false
	}
}
