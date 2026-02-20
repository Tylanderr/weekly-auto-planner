package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/magiconair/properties"
	"github.com/yosssi/gohtml"

	"net/smtp"
	"os"
	"slices"

	"github.com/tylander732/weeklyAutoPlanner/pkg/consts"
	"github.com/tylander732/weeklyAutoPlanner/pkg/html"
	"github.com/tylander732/weeklyAutoPlanner/pkg/model"
)

var propertiesFile = "./resources/app.properties"
var username string
var password string

var sendEmailFlag bool = false

var groceryCategories = []consts.GroceryCategory{
	consts.Produce,
	consts.MeatAndPoultry,
	consts.Seafood,
	consts.Dairy,
	consts.Bakery,
	consts.FrozenFoods,
	consts.PantryStaples,
	consts.Beverages,
	consts.Snacks,
	consts.HouseholdGoods,
	consts.PersonalCare,
	consts.InternationalFoods,
	consts.Deli,
}

func main() {
	readProperties()
}

func sendEmail(emailString string, receiver string) {
	// smtp server configuration
	smtpHost := "smtp.gmail.com"
	server := smtpHost + ":587"

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	auth := smtp.PlainAuth("", username, password, smtpHost)

	subject := "Subject: Auto Emailer Test\n"

	msg := []byte(subject + mime + emailString)

	err := smtp.SendMail(server, auth, username, []string{receiver}, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent successfully")
}

func readProperties() {
	p := properties.MustLoadFile(propertiesFile, properties.UTF8)
	username, _ = p.Get("username")
	password, _ = p.Get("password")
}
