package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChannel := make(chan kafka.Event)
	producer := NewKafkaProducer()

	Publish("Heavy Metal", "desafio-2", producer, deliveryChannel)

	DeliveryReport(deliveryChannel)
}