package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func consumer(broker string, topics []string, group string) {
	log.Println("Consumer (Kafka) started!")

	// Setting up Consumer (Kafka) config
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          group,
		"auto.offset.reset": "latest",
	})

	if err != nil {
		panic(err)
	}

	// Subscribe to topics
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			log.Println("New Request from Kafka")
			log.Printf("Message consumed on %s: %s\n", msg.TopicPartition, string(msg.Value))
			headerRes := msg.Headers
			headerVal := string(headerRes[0].Value)

			res := consume{
				header: headerVal,
				body:   string(msg.Value),
			}

			fmt.Println(res.header)
			// Send any consumed event to consumerChan
			consumerChan <- res
			consumerChan2 <- res
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
