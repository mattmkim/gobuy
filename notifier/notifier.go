package notifier

import (
	"fmt"
	"log"
	"net/smtp"
)

func Notify(config map[string]string, service string, message string) {
	to := config["notifier.to"]
	from := config["notifier.from"]
	password := config["notifier.password"]
	msg := fmt.Sprintf("Subject: Alert from %s Service\n\n %s", service, message)
	status := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, password, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if status != nil {
		log.Printf("Error from SMTP Server: %s", status)
	}
	log.Print("Email Sent Successfully")
}
