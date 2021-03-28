package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func producer(broker string, topic string, message <-chan consume) {
	log.Println("Producer started!")
	rand.Seed(time.Now().UnixNano())

	// Setting up kafka message to get ready to be produced
	// message to be produced
	msg := <-message

	nice := rand.Intn(100)
	time.Sleep(time.Duration(nice) * time.Millisecond)
	fmt.Println(strconv.Itoa(nice) + "timing")
	// Setting up Consumer (Kafka) config
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// Run go routine for produce available event to Kafka
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Produce failed: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Produced message to %v. Message: %s (Header: %s)\n", ev.TopicPartition, ev.Value, ev.Headers)
				}
			}
		}
	}()

	// header for the message
	header := map[string]string{
		"key":   "uniqueValue",
		"value": msg.header,
	}
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg.body),
		Headers:        []kafka.Header{{Key: header["key"], Value: []byte(header["value"])}},
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(3 * 1000)
	log.Println("Producer closing!")
}
