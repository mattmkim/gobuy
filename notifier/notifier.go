package notifier

import (
	"fmt"
	"log"
	"net/smtp"
)

func Notify(service string, message string) {
	// TODO: create secrets file for email auth
	to := ""
	from := ""
	password := ""
	msg := fmt.Sprintf("Subject: Alert form %s Service\n\n %s", service, message)
	status := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, password, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if status != nil {
		log.Printf("Error from SMTP Server: %s", status)
	}
	log.Print("Email Sent Successfully")
}
