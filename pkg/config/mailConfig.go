package config

import (
	"fmt"
	"net/smtp"
	"strconv"
)

type MailConfig interface {
	SendMail(cfg Config, to string, message []byte) error
}

type mailConfig struct{}

func NewMailConfig() MailConfig {
	return &mailConfig{}

}

func (c *mailConfig) SendMail(cfg Config, to string, message []byte) error {
	//	log.Println("Email is to senf message:", to)
	fmt.Println(to, message)
	userName := cfg.SMTPUSERNAME
	password := cfg.SMTPHTTPPASSWORD
	smtpHost := cfg.SMTPHOST
	smtpPort := cfg.SMTPPORT
	auth := smtp.PlainAuth("", userName, password, smtpHost)

	fmt.Printf("\n\nTO%v\n\n", to)
	
	num, _ := strconv.Atoi(smtpPort)
	//err := smtp.SendMail(smtpHost+":"+smtpPort, auth, userName, []string{to}, message)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, num), auth, userName, []string{to}, []byte(message))

	fmt.Printf("\n\nerror%v\n\n", err)
	//sending email
	return err
}
