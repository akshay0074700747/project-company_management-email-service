package main

import (
	"log"
	"os"

	"github.com/akshay0074700747/email-service/email"
	"github.com/akshay0074700747/email-service/kafka"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	smtp_email := os.Getenv("EMAIL")
	smtp_password := os.Getenv("PASSWORD")
	smtp_server := "smtp.gmail.com"
	smtp_port := "587"

	mailer := email.NewMailer(smtp_server, smtp_port, smtp_email, smtp_password)
	kafka.Getmailer(mailer)

	kafka.StartServing()
}
