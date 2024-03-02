package services

// Ethereal Email
// Email: kurt.heaney@ethereal.email
// Pass: 2g8QsnyNdt2zhChpQ8

import (
	"games-email/models"

	"github.com/wneessen/go-mail"
)

// TODO: change emails to dynamic ones once fiture out fix
func SendUserEmail(user *models.UserEmail) error {
	message := mail.NewMsg()	
	message.Subject("User Password Update")
	message.SetBodyString(mail.TypeTextPlain, "Your Password has been updated")
	err := sendEmail(message, *user.Email)
	if err != nil {
		return err
	}
	return nil
}
