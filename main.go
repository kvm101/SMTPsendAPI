package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func sendEmail(to []string, subject string, body string) error {
	auth := smtp.PlainAuth("",
		os.Getenv("FROM_EMAIL"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("EMAIL_SMTP"),
	)

	msg := []byte("Subject: " + subject + "\n" + body)

	return smtp.SendMail(
		os.Getenv("EMAIL_SMTP_ADDR"),
		auth,
		os.Getenv("FROM_EMAIL"),
		to,
		msg,
	)
}

func main() {
	godotenv.Load()

	err := sendEmail([]string{"mail@gmail.com", "mail@gmail.com"}, "---...---", "---...---")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Email was send!")
	}
}
