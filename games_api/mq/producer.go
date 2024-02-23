package mq

import (
	"fmt"
	"github.com/IBM/sarama"
)

func CreateMessage(topic string, content string) {
	brokerList := []string{"kafka:9092"}
	producer, err := sarama.NewSyncProducer(brokerList, nil)
	if err != nil {
		fmt.Println("Producer creation failed", err)
		return
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(content),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		fmt.Println("Send message failed", err)
		return
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", message.Topic, partition, offset)
}
