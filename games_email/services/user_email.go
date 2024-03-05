package services

// Ethereal Email
// Email: kurt.heaney@ethereal.email
// Pass: 2g8QsnyNdt2zhChpQ8

import (
	"games-email/models"

	"github.com/wneessen/go-mail"
)

func SendUserEmail(user *models.UserEmail) error {
	message := mail.NewMsg()
	message.Subject("Password Updated")
	message.SetBodyString(mail.TypeTextPlain, "Your password has been updated!")
	// err := sendEmail(message, *user.Email)
	err := sendEmail(message, "kurt.heaney@ethereal.email")
	if err != nil {
		return err
	}
	return nil
}
