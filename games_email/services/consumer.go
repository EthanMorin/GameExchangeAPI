package services

import (
	"context"
	"fmt"
	// "time"

	kafka "github.com/segmentio/kafka-go"
)

func StartListener() {
	topic := "email"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, partition)
	if err != nil {
		fmt.Println("Error connecting to kafka", err)
	}
	// conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6)

	b := make([]byte, 10e3)
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		fmt.Println("Error closing batch", err)
	}

	if err := conn.Close(); err != nil {
		fmt.Println("Error closing connection", err)
	}
}