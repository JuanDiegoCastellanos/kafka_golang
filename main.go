package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	topic := "HVSE"
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "foo",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)

	}
	go func() {
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
			"group.id":          "foo",
			"auto.offset.reset": "smallest"})

		if err != nil {
			log.Fatal(err)
		}
		err = consumer.Subscribe(topic, nil)

		if err != nil {
			log.Fatal(err)
		}

		for {
			ev := consumer.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("consumed message from the queue:  %s \n", string(e.Value))
			case *kafka.Error:
				fmt.Printf("%s \n", e)
			}
		}

	}()

	deliverych := make(chan kafka.Event, 10000)

	for {

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte("FOO")},
			deliverych,
		)
		if err != nil {
			log.Fatal(err)
		}

		<-deliverych
		time.Sleep(time.Second * 3)
	}
}
