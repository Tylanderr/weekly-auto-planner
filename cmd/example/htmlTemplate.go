package sample

import (
    "bytes"
    "html/template"
    "net/smtp"
)

type EmailData struct {
    Name  string
    Email string
}

func main() {
    // Sample email data
    data := EmailData{
        Name:  "John Doe",
        Email: "john@example.com",
    }

    // Execute the HTML template
    emailBody, err := executeTemplate("email_template.html", data)
    if err != nil {
        panic(err)
    }

    // Send email
    sendEmail("sender@example.com", "recipient@example.com", "Sample Subject", emailBody)
}

func executeTemplate(templateFile string, data EmailData) (string, error) {
    // Parse the template file
    tmpl, err := template.ParseFiles(templateFile)
    if err != nil {
        return "", err
    }

    // Execute the template with the provided data
    var tpl bytes.Buffer
    err = tmpl.Execute(&tpl, data)
    if err != nil {
        return "", err
    }

    return tpl.String(), nil
}

func sendEmail(from, to, subject, body string) error {
    // SMTP server configuration
    smtpServer := "smtp.example.com"
    smtpPort := "587"
    smtpUsername := "your_smtp_username"
    smtpPassword := "your_smtp_password"

    // Authentication
    auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)

    // Compose email message
    message := "From: " + from + "\r\n" +
        "To: " + to + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "Content-Type: text/html; charset=UTF-8\r\n" +
        "\r\n" +
        body

    // Send email
    err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, []byte(message))
    if err != nil {
        return err
    }

    return nil
}

