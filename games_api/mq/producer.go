// Reference for how to use kafka-go
// https://github.com/segmentio/kafka-go
package mq

import (
	"context"
	"fmt"

	kafka "github.com/segmentio/kafka-go"
)

func connectToKafka() *kafka.Conn {
	topic := "email"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		fmt.Println("Error connecting to kafka", err)
	}
	return conn
}

func SendEmail(email string, message string) {
	conn := connectToKafka()
	conn.WriteMessages(
		kafka.Message{Value: []byte(email)},
		kafka.Message{Value: []byte(message)},
	)
	conn.Close()
}