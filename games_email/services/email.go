package services

import (
	"log"

	"github.com/wneessen/go-mail"
)

func sendEmail(message *mail.Msg) {
	c, err := mail.NewClient("smtp.ethereal.email", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername("kurt.heaney@ethereal.email"), mail.WithPassword("2g8QsnyNdt2zhChpQ8"))
	if err != nil {
		log.Fatalf("failed to create mail client: {%s}", err)
	}
	if err := c.DialAndSend(message); err != nil {
		log.Fatalf("failed to send mail: {%s}", err)
	}
}