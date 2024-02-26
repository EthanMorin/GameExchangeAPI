package services

// Ethereal Email
// Email: kurt.heaney@ethereal.email
// Pass: 2g8QsnyNdt2zhChpQ8

import (
	"log"

	"github.com/wneessen/go-mail"
)

func SendEmail(email string, subject string, body string) {
	message := mail.NewMsg()
	if err := message.From("kurt.heaney@ethereal.email"); err != nil {
		log.Fatalf("Failed to set sender: {%s}", err)
	}
	if err := message.To(email); err != nil {
		log.Fatalf("Failed to set recipient: {%s}", err)
	}
	message.Subject(subject)
	message.SetBodyString(mail.TypeTextPlain, body)

	c, err := mail.NewClient("smtp.ethereal.email", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername("kurt.heaney@ethereal.email"), mail.WithPassword("2g8QsnyNdt2zhChpQ8"))
	if err != nil {
		log.Fatalf("failed to create mail client: {%s}", err)
	}
	if err := c.DialAndSend(message); err != nil {
		log.Fatalf("failed to send mail: {%s}", err)
	}
}
