package services

// Ethereal Email
// Email: kurt.heaney@ethereal.email
// Pass: 2g8QsnyNdt2zhChpQ8

import (
	"games-email/models"
	"log"

	"github.com/wneessen/go-mail"
)

func SendUserEmail(user *models.UserEmail) {
	message := mail.NewMsg()
	if err := message.From("kurt.heaney@ethereal.email"); err != nil {
		log.Fatalf("Failed to set sender: {%s}", err)
	}
	
	// Why does this work when i type it out but not pull from the model?
	if err := message.To("kurt.heaney@ethereal.email"); err != nil {
		// if err := message.To(*user.Email); err != nil {
		log.Fatalf("Failed to set recipient: {%s}", err)
	}

	message.Subject("User Password Update")
	message.SetBodyString(mail.TypeTextPlain, "Your Password has been updated")

	c, err := mail.NewClient("smtp.ethereal.email", mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername("kurt.heaney@ethereal.email"), mail.WithPassword("2g8QsnyNdt2zhChpQ8"))
	if err != nil {
		log.Fatalf("failed to create mail client: {%s}", err)
	}
	if err := c.DialAndSend(message); err != nil {
		log.Fatalf("failed to send mail: {%s}", err)
	}
}
