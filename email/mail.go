package email

import (
	"fmt"
	"log"
	"net/smtp"
)

type SMTPConfig struct {
	SMTPServer   string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

func NewMailer(smtpServer, smtpPort, smtpUsername, smtpPassword string) *SMTPConfig {
	return &SMTPConfig{
		SMTPServer:   smtpServer,
		SMTPPassword: smtpPassword,
		SMTPUsername: smtpUsername,
		SMTPPort:     smtpPort,
	}
}

func (config *SMTPConfig) SendMessage(recipient string, message string) error {

	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPServer)

	err := smtp.SendMail(
		"",
		auth,
		config.SMTPUsername,
		[]string{recipient},
		[]byte(message),
	)

	if err != nil {
		log.Printf("Error sending email: %v\n", err)
		return err
	}

	fmt.Println("Email sent successfully!")
	return nil
}
