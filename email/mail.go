package email

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
)

type SMTPConfig struct {
	SMTPServer   string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

func NewMailer(smtpServer, smtpPort, smtpUsername, smtpPassword string) *SMTPConfig {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return &SMTPConfig{
		SMTPServer:   smtpServer,
		SMTPPassword: smtpPassword,
		SMTPUsername: smtpUsername,
		SMTPPort:     smtpPort,
	}
}

func (config *SMTPConfig) SendMessage(recipient string, message string) error {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	auth := smtp.PlainAuth("", config.SMTPUsername, config.SMTPPassword, config.SMTPServer)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	err := smtp.SendMail(
		config.SMTPServer+":"+config.SMTPPort,
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
