package main

import (
	"context"
	"encoding/json"
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
            var user models.UserEmail
            _ = json.Unmarshal(msg.Value, &user)
            err := services.SendUserEmail(&user)
            if err != nil {
                log.Fatalf("Couldnt send email: {%s}", err)
            }
        case "exchange":
            var offer models.EchangeEmails
            _ = json.Unmarshal(msg.Value, &offer)
            err := services.SendExchangeEmail(string(msg.Key), &offer)
            if err != nil {
                log.Fatalf("Couldnt send email: {%s}", err)
            }
        }
        sess.MarkMessage(msg, "")
    }
    return nil
}

func main() {
    brokers := []string{"kafka:9092"}
    topics := []string{"user", "exchange"}

    config := sarama.NewConfig()
    config.Consumer.Return.Errors = true

    group, err := sarama.NewConsumerGroup(brokers, "email-group", config)
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