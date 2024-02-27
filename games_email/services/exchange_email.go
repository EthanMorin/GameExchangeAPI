package services

// Ethereal Email
// Email: kurt.heaney@ethereal.email
// Pass: 2g8QsnyNdt2zhChpQ8

import (
	"games-email/models"
	"log"

	"github.com/wneessen/go-mail"
)

// TODO: change emails to dynamic ones once fiture out fix
func SendExchangeEmail(topic string, exchange *models.EchangeEmails) {
	message := mail.NewMsg()
	switch topic {
	case "exchange_created":
		message.Subject("Offer Created")
		message.SetBodyString(mail.TypeTextPlain, "You have a pending offer")
		// if err := message.To(*exchange.TraderEmail); err != nil {
		if err := message.To("kurt.heaney@ethereal.email"); err != nil {
			log.Fatalf("Failed to set recipient: {%s}", err)
		}
		sendEmail(message)
		message.SetBodyString(mail.TypeTextPlain, "Your offer has been created")
	case "exchange_accepted":
		message.Subject("Offer Accepted")
		message.SetBodyString(mail.TypeTextPlain, "Your offer has been accepted")
	case "exchange_rejected":
		message.Subject("Offer Rejected")
		message.SetBodyString(mail.TypeTextPlain, "Your offer has been rejected")
	}
	if err := message.From("kurt.heaney@ethereal.email"); err != nil {
		log.Fatalf("Failed to set sender: {%s}", err)
	}
	// if err := message.To(*exchange.TradeeEmail); err != nil {
	if err := message.To("kurt.heaney@ethereal.email"); err != nil {
		log.Fatalf("Failed to set recipient: {%s}", err)
	}
	sendEmail(message)
}
