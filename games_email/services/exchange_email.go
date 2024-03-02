package services

// Ethereal Email
// Email: kurt.heaney@ethereal.email
// Pass: 2g8QsnyNdt2zhChpQ8

import (
	"errors"
	"games-email/models"

	"github.com/wneessen/go-mail"
)

func SendExchangeEmail(topic string, exchange *models.EchangeEmails) error {
	message := mail.NewMsg()

	switch topic {
	case "exchange_created":
		message.Subject("Offer Created")
		message.SetBodyString(mail.TypeTextPlain, "You have a pending offer!")
		err := sendEmail(message, *exchange.TraderEmail)
		if err != nil {
			return err
		}
		message.SetBodyString(mail.TypeTextPlain, "Your offer has been created!")
	case "exchange_accepted":
		message.Subject("Offer Accepted")
		message.SetBodyString(mail.TypeTextPlain, "Your offer has been accepted!")
	case "exchange_rejected":
		message.Subject("Offer Rejected")
		message.SetBodyString(mail.TypeTextPlain, "Your offer has been rejected!")
	default:
		return errors.New("topic doesnt match")
	}
	err := sendEmail(message, *exchange.TradeeEmail)
	if err != nil {
		return err
	}
	return nil
}
