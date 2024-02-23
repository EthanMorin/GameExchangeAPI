package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"os/signal"
)

func main() {
	brokers := []string{"kafka:9092"}
	topics := []string{"users", "trades"}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Panicf("Error creating consumer: %v", err)
	}
	defer consumer.Close()

	messages := make(chan *sarama.ConsumerMessage)

	for _, topic := range topics {
		partitions, err := consumer.Partitions(topic)
		if err != nil {
			log.Panicf("Error getting partitions for topic %s: %v", topic, err)
		}

		for _, partition := range partitions {
			pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
			if err != nil {
				log.Panicf("Error consuming partition %d of topic %s: %v", partition, topic, err)
			}
			go func(pc sarama.PartitionConsumer) {
				for msg := range pc.Messages() {
					messages <- msg
				}
			}(pc)
		}
	}

	go func() {
		for msg := range messages {
			fmt.Printf("Received message: topic=%s, partition=%d, offset=%d, value=%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	<-sigterm
}
