package main

import "github.com/confluentinc/confluent-kafka-go/kafka"

func main() {
	deliveryChannel := make(chan kafka.Event)
	processor := NewKafkaProcessor(deliveryChannel)
	processor.Consume()
}