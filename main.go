package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "something",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)

	}
	deliverych := make(chan kafka.Event, 10000)
	topic := "HVSE"
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("FOO")},
		deliverych,
	)
	if err != nil {
		log.Fatal(err)
	}

	e := <-deliverych
	fmt.Printf("%+v\n", e.String())
	fmt.Printf("%+v\n", p)
}
