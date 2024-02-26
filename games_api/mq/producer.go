package mq

import (
	"encoding/json"
	"games-api/models"

	"github.com/IBM/sarama"
)

var producer, _ = sarama.NewSyncProducer([]string{"kafka9092"}, nil)

func UpdateUserPass(user *models.User) {
	userJson, _ := json.Marshal(user)
	msg := &sarama.ProducerMessage{
		Topic: "user",
		Key:   sarama.StringEncoder("password_update"),
		Value: sarama.StringEncoder(userJson),
	}
	producer.SendMessage(msg)
}

func CreateExchange(exchange *models.Exchange) {
	exchangeJson, _ := json.Marshal(exchange)
	msg := &sarama.ProducerMessage{
		Topic: "user",
		Key:   sarama.StringEncoder("exchange_created"),
		Value: sarama.StringEncoder(exchangeJson),
	}
	producer.SendMessage(msg)
}

func UpdateExchangeAccepted(exchange *models.Exchange) {
	exchangeJson, _ := json.Marshal(exchange)
	msg := &sarama.ProducerMessage{
		Topic: "user",
		Key:   sarama.StringEncoder("exchange_accepted"),
		Value: sarama.StringEncoder(exchangeJson),
	}
	producer.SendMessage(msg)
}

func UpdateExchangeRejected(exchange *models.Exchange) {
	exchangeJson, _ := json.Marshal(exchange)
	msg := &sarama.ProducerMessage{
		Topic: "user",
		Key:   sarama.StringEncoder("exchange_rejected"),
		Value: sarama.StringEncoder(exchangeJson),
	}
	producer.SendMessage(msg)
}
