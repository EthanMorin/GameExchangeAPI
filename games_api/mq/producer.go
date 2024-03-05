package mq

import (
	"encoding/json"
	"errors"
	"games-api/models"
	"log"

	"github.com/IBM/sarama"
)


func UpdateUserPass(user *models.User) {
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("couldnt create producer: {%s}", err)
	}
	userJson, err := json.Marshal(&user)
	if err != nil {
		log.Fatalf("couldnt Marshal json: {%s}", err)
	}
	msg := &sarama.ProducerMessage{
		Topic: "user",
		Key:   sarama.StringEncoder("password_update"),
		Value: sarama.StringEncoder(userJson),
	}
	producer.SendMessage(msg)
}

func CreateExchange(exchange *models.Exchange) {
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("couldnt create producer: {%s}", err)
	}
	exchangeJson, err := json.Marshal(exchange)
	if err != nil {
		log.Fatalf("couldnt Marshal json: {%s}", err)
	}
	msg := &sarama.ProducerMessage{
		Topic: "exchange",
		Key:   sarama.StringEncoder("exchange_created"),
		Value: sarama.StringEncoder(exchangeJson),
	}
	producer.SendMessage(msg)
}

func UpdateExchange(exchange *models.Exchange) error {
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		return errors.New("couldn't create producer: " + err.Error())
	}
	exchangeJson, err := json.Marshal(exchange)
	if err != nil {
		return errors.New("couldn't Marshal json: " + err.Error())
	}

	var msg *sarama.ProducerMessage
	if *exchange.Status == "accepted" {
			msg = &sarama.ProducerMessage{
					Topic: "exchange",
					Key:   sarama.StringEncoder("exchange_accepted"),
					Value: sarama.StringEncoder(exchangeJson),
			}
	} else {
			msg = &sarama.ProducerMessage{
					Topic: "exchange",
					Key:   sarama.StringEncoder("exchange_rejected"),
					Value: sarama.StringEncoder(exchangeJson),
			}
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}

	return nil
}