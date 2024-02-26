package mq

import (
	"encoding/json"
	"games-api/models"
	"log"

	"github.com/IBM/sarama"
)

func UpdateUserPass(user *models.User){
	producer, err := sarama.NewSyncProducer([]string{"kafka9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}

	userJson, _ := json.Marshal(user)
	msg := &sarama.ProducerMessage{
		Topic: "user",
		Key: sarama.StringEncoder("password_update"),
		Value: sarama.StringEncoder(userJson),
	}

	producer.SendMessage(msg)
}

