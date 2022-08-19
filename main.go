package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	var email string = os.Getenv("EMAIL")
	var mailPassword string = os.Getenv("PASSWORD")

	//Mail information
	from := email
	password := mailPassword

	//SMTP information
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	t, _ := template.ParseFiles("template/index.html")

	var body bytes.Buffer

	// Reciever
	to := []string{"ferrivas127@gmail.com"}

	// Message
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Puneet Singh",
		Message: "This is a test message in a HTML template",
	})

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("All good!")
}
