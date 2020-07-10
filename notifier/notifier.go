package notifier

import (
	"fmt"
	"net/smtp"

	log "github.com/sirupsen/logrus"
)

func Notify(config map[string]string, service string, message string) {
	msg := fmt.Sprintf("Subject: Alert from %s Service\n\n %s", service, message)
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", config["notifier.from"], config["notifier.password"], "smtp.gmail.com"), config["notifier.from"], []string{config["notifier.to"]}, []byte(msg))

	if err != nil {
		log.Printf("Error from SMTP Server: %s", err)
	}
	log.Print("Email Sent Successfully")
}
