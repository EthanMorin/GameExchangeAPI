package services

// Ethereal Email
// Email: kurt.heaney@ethereal.email
// Pass: 2g8QsnyNdt2zhChpQ8

import (
	"games-email/data"
	"games-email/models"
	"log"

	"github.com/wneessen/go-mail"
)

func SendExchangeEmail(topic string, exchange *models.EchangeIds) {
	message := mail.NewMsg()
	if err := message.From("kurt.heaney@ethereal.email"); err != nil {
		log.Fatalf("Failed to set sender: {%s}", err)
	}

	var traderEmail string
	var tradeeEmail string
	switch topic {
	case "exchange_created":
		message.Subject("Offer Created")
		message.SetBodyString(mail.TypeTextPlain, "Your offer has been created")
		traderEmail, _ = data.GetUserEmail(*exchange.Traderid)
		tradeeEmail, _ = data.GetUserEmail(*exchange.Tradeeid)
	case "exchange_accepted":
		message.Subject("Offer Accepted")
		message.SetBodyString(mail.TypeTextPlain, "Your offer has been accepted")
		traderEmail = ""
		tradeeEmail, _ = data.GetUserEmail(*exchange.Tradeeid)
	case "exchange_rejected":
		message.Subject("Offer Rejected")
		message.SetBodyString(mail.TypeTextPlain, "Your offer has been rejected")
		traderEmail = ""
		tradeeEmail, _ = data.GetUserEmail(*exchange.Tradeeid)
	}
	if err := message.To(tradeeEmail); err != nil {
		log.Fatalf("Failed to set recipient: {%s}", err)
	}
	sendEmail(message)
	if traderEmail != "" {
		message.Subject("Offer Created")
		message.SetBodyString(mail.TypeTextPlain, "You have a pending offer")
		if err := message.To(traderEmail); err != nil {
			log.Fatalf("Failed to set recipient: {%s}", err)
		}
		sendEmail(message)
	}
}
