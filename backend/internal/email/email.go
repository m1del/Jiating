package email

import (
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(name, fromEmail, subject, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_USER"))
	m.SetHeader("To", os.Getenv("EMAIL_USER"))
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", "Message from "+name+" ("+fromEmail+"):\n\n"+message)

	// setup auth info
	d := gomail.NewDialer(
		os.Getenv("EMAIL_HOST"), 587,
		os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"),
	)

	// send the email
	return d.DialAndSend(m)
}
