package main

import (
	"log"
	"net/http"
	"os"
	"sync"
)

var billerChanReq = make(chan string)
var billerChanRes = make(chan string)
var producerChan = make(chan string)
var consumerChan = make(chan string)

func main() {
	// Service log setup
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Found error in log ", err)
	}
	log.SetOutput(file)

	// Service setup
	router := server()
	go func() {
		log.Fatal(http.ListenAndServe(":6020", router))
	}()

	//create wait group
	var wg sync.WaitGroup

	//consumer routine
	go consumer()
	go respIso()

	//producer routine
	for {
		select {
		case x := <-billerChanRes:
			log.Println("New request in billerChan is ready to produce")
			wg.Add(1)
			go producer(&wg, producerChan)
			producerChan <- x
			wg.Wait()
		default:
			continue
		}
	}

}
