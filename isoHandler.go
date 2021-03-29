package main

import (
	"fmt"
	"log"

	"github.com/mofax/iso8583"
)

// Handler to new consumed request in consumerChan and send new response to billerChan
func requestHandler() {

	// loop for checking if there is any new request from Consumer (Kafka) that has been sent to consumerChan
	for {
		select {
		// execute if there is a new request in consumerChan
		case newRequest := <-consumerChan2:

			// Send new request to `Biller` and get response that ready to produce
			msg := newRequest.body
			getResponse(msg)
			// Send new response to billerChan
			//billerChan <- isoParsed
			//fmt.Println(isoParsed)
			// Done with requestHandler
			log.Println(newRequest)

		// keep looping if there is none new request
		default:
			continue
		}
	}

}

func getResponse(message string) {
	//var response iso8583

	isoContent := message[4:]

	isoStruct := iso8583.NewISOStruct("spec1987.yml", false)
	msg, err := isoStruct.Parse(isoContent)
	if err != nil {
		log.Println(err)
	}

	//var isoParsed iso8583.IsoStruct
	jsonIso := getEpayRinstis(msg)

	responseFromBiller := sendJsonToBiller(jsonIso, "rintis")

	fmt.Println(responseFromBiller)
}
