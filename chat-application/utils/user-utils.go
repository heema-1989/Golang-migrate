package utils

import (
	"bytes"
	"chat-application/models"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"strconv"
)

func GetSMTPCredentials() (from, username, password string, port int) {
	from = os.Getenv("SMTP_FROM")
	username = os.Getenv("SMTP_USERNAME")
	password = os.Getenv("SMTP_PASSWORD")
	port, _ = strconv.Atoi(os.Getenv("SMTP_PORT"))
	return from, username, password, port
}

func TemplateParser(user models.Token) (message string) {
	var (
		temp       *template.Template
		parsingErr error
	)
	temp, parsingErr = temp.ParseFiles("/home/heema/GolandProjects/golang-migrate/chat-application/views/default/verification-email.html")
	CheckError(parsingErr, "Error parsing email send temp")
	buff := new(bytes.Buffer)
	executeErr := temp.Execute(buff, user)
	CheckError(executeErr, "Error executing email temp")
	message = buff.String()
	return message
}

func SendMail(emailTo string, user models.Token) {
	var (
		emailFrom, username, password string
		port                          int
	)
	emailFrom, username, password, port = GetSMTPCredentials()
	email := gomail.NewMessage()
	email.SetHeader("From", emailFrom)
	email.SetHeader("To", emailTo)
	email.SetHeader("Subject", "Please verify your email")
	email.SetBody("text/html", TemplateParser(user))
	dialer := gomail.NewDialer("smtp.gmail.com", port, username, password)
	dialErr := dialer.DialAndSend(email)
	CheckError(dialErr, "Error sending email")
}
