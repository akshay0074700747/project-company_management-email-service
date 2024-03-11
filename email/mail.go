package email

import (
	"fmt"
	"net/smtp"
)


type SMTPConfig struct {
	SMTPServer   string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

func NewMailer(smtpServer,smtpPort,smtpUsername,smtpPassword string) *SMTPConfig {
	return &SMTPConfig{
		SMTPServer: smtpServer,
		SMTPPassword: smtpPassword,
		SMTPUsername: smtpUsername,
		SMTPPort: smtpPort,
	}
}


func (config *SMTPConfig) SendMessage(reciever string,message string) error {
	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPServer)

	if err := smtp.SendMail(config.SMTPServer+":"+config.SMTPPort, auth, config.SMTPUsername, []string{reciever}, []byte(message)); err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("email sent...")

	return nil
}
