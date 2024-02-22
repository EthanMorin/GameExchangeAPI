package mq

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	Producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(brokers string, topic string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		Producer: p,
		topic:    topic,
	}, nil
}

func (kp *KafkaProducer) ProduceMessage(key, message string) error {
	deliveryChan := make(chan kafka.Event)
	err := kp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(message),
	}, deliveryChan)
	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}

func (kp *KafkaProducer) Close() {
	kp.Producer.Close()
}

