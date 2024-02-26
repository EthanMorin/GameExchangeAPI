package main

import (
	"context"
	"encoding/json"
	"games-email/data"
	"games-email/models"
	"games-email/services"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/IBM/sarama"
)

type groupHandler struct{}

func (groupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (groupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h groupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		switch msg.Topic {
		case "user":
			var user models.UserIdEmail
			_ = json.Unmarshal(msg.Value, &user)
			services.SendUserEmail(string(msg.Topic), &user)
		case "exchange":
			var exchange models.EchangeIds
			_ = json.Unmarshal(msg.Value, &exchange)
			services.SendExchangeEmail(string(msg.Key), &exchange)
		}
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	brokers := []string{"kafka:9092"}
	topics := []string{"user", "exchange"}
	err := data.NewDB()
	if err != nil {
		log.Fatalf("Couldnt establish connection to DB: {%s}", err)
	}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	group, err := sarama.NewConsumerGroup(brokers, "example-group", config)
	if err != nil {
		log.Panicf("Error creating consumer group: %v", err)
	}
	defer func() { _ = group.Close() }()

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := group.Consume(ctx, topics, groupHandler{}); err != nil {
				panic(err)
			}
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	<-sigterm

	cancel()
	wg.Wait()
}
