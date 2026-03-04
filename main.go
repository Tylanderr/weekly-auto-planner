package main

import (
	"fmt"
	"os"

	"net/smtp"
)

var username = os.Getenv("HOP_EMAIL_USERNAME")
var password = os.Getenv("HOP_EMAIL_PASSWORD")

var sendEmailFlag bool = false

func main() {

}

func sendEmail(emailString string, receiver string) {
	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	server := smtpHost + ":587"

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	auth := smtp.PlainAuth("", username, password, smtpHost)

	subject := "Subject: Auto Emailer Test\n"

	msg := []byte(subject + mime + emailString)

	if sendEmailFlag {
		err := smtp.SendMail(server, auth, username, []string{receiver}, msg)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Email sent successfully")
}
