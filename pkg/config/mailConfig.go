package config

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
)

type MailConfig interface {
	SendMail(cfg Config, to string, message string) error
}

type mailConfig struct{}

func NewMailConfig() MailConfig {
	return &mailConfig{}

}

func (c *mailConfig) SendMail(cfg Config, to string, message string) error {
	fmt.Println(to, message, "kahjsdfikhdsfkjsdhnfijusdhfujihasfiksjhsjfsdojikfghhsjvgahjsklhasbdfjgb\n\n\n\n\n ")
	log.Println("Email is to senf message:", to)
	fmt.Println(to, message)
	userName := cfg.SMTPUSERNAME
	password := cfg.SMTPHTTPPASSWORD
	smtpHost := cfg.SMTPHOST
	smtpPort := cfg.SMTPPORT
	auth := smtp.PlainAuth("", userName, password, smtpHost)
	headers := make(map[string]string)
	headers["Subject"] = "ecommerce"
	headers["From"] = "userName"
	var msg bytes.Buffer
	for k, v := range headers {
		msg.WriteString(k + ": " + v + "\n")

	}
	msg.WriteString("\r \n")
	msg.WriteString(message)

	//sending email
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, userName, []string{to}, msg.Bytes())
}
