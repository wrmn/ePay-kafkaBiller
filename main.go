package main

import (
	"log"
	"os"
)

type consume struct {
	header string
	body   string
}

var (
	billerChan    = make(chan string)  // channel for send-receive data from-to `Biller`
	producerChan  = make(chan consume) // channel for receive data from channelChan and send data to `Producer (Kafka)`
	consumerChan  = make(chan consume) // channel for receive data from `Consumer (Kafka)` and send data to channelChan
	consumerChan2 = make(chan consume) // channel for receive data from `Consumer (Kafka)` and send data to channelChan
)

func main() {
	// Setting up log file
	// set permission to read/write log file
	// read/write to existing log file, if there is none it will create new log file
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Found error in log ", err)
	}
	log.SetOutput(file)

	// ChannelKafka started
	log.Println("Service Started!")

	// Setting up HTTP Listener and Handler
	// router will handle any request at any endpoint available in server()

	// WaitGroup for make sure that an event is produced to Kafka

	// Get config for Kafka Producer and Consumer
	broker, producerTopic, consumerTopics, groups := configKafka()

	// Run Consumer (Kafka)
	go consumer(broker, consumerTopics, groups)

	go requestHandler()

	// Run Goroutine for request-response data from-to `Biller`
	//go requestHandler()

	// loop for checking if there is any new response from `Biller` that has been sent to channelChan
	for {
		select {
		// execute if there is a new response in channelChan
		case newResponse := <-consumerChan:
			log.Println("New response from `Biller` is ready to produce to Kafka")

			// Add WaitGroup counter to wait until Producer (Kafka) finish producing an event

			// Run Producer (Kafka)
			go producer(broker, producerTopic, producerChan)

			// Send new response to producerChan, then produce the new response to Kafka
			producerChan <- newResponse

			// Waiting until Producer (Kafka) finish producing event

		// keep looping if there is none new response
		default:
			continue
		}
	}
}
