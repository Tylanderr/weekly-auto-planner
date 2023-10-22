package main

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"
)

type User struct {
    Email string `json:"email"`
    Meal Meal `json:"meal"`
}

type Meal struct {
    Name string `json:"name"`
    Ingrediants []string `json:"ingrediants"`
}

func main() {
    readJsonFile()
}

func readJsonFile() {
    var firstAttempt Meal

    contents, err := os.ReadFile("./resources/userList.json")
    if err != nil {
        fmt.Println("error reading the contents of the file: ", err)
        return
    }
    // fmt.Println(string(contents))

    err2 := json.Unmarshal(contents, &firstAttempt)

    if err2 != nil {
        fmt.Println("unable to convert contents to struct", err2)
        return
    }

    fmt.Println(firstAttempt)
}


func sendEmail() {
    from := ""
    password := ""

    to := []string{""}

    // smtp server configuration
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"


    message := []byte("This is a test email message")

    auth := smtp.PlainAuth("", from, password, smtpHost)

    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Email sent successfully")
}


