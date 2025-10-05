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
