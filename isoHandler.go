package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mofax/iso8583"
	"github.com/rivo/uniseg"
)

// Handler to new consumed request in consumerChan and send new response to billerChan
func requestHandler() {

	// loop for checking if there is any new request from Consumer (Kafka) that has been sent to consumerChan
	for {
		select {
		// execute if there is a new request in consumerChan
		case newRequest := <-consumerChan:

			// Send new request to `Biller` and get response that ready to produce
			msg := newRequest
			isoParsed := getResponse(msg)

			// Send new response to billerChan
			billerChan <- isoParsed

			// Done with requestHandler
			log.Println("New request handled")

		// keep looping if there is none new request
		default:
			continue
		}
	}

}

// Return response from `Biller` in ISO8583 Format
func getResponse(message string) (isoResponse string) {

	var response Iso8583
	data := message[4:]

	// Parse new ISO8583 message to ISO Struct
	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)
	msg, err := isoStruct.Parse(data)
	if err != nil {
		log.Println(err)
	}

	var isoParsed iso8583.IsoStruct

	// Check processing code and send request to appropriate `Biller` endpoints
	pcode := msg.Elements.GetElements()[3]
	switch pcode {
	// Process PPOB Inquiry request
	case "380001":
		// Convert ISO message to JSON format
		jsonIso := getJsonPPOBInquiry(msg)

		// Send JSON data to Biller
		serverResp := responseJsonPPOBInquiry(jsonIso)

		// Convert response from JSON data to ISO8583 format
		isoParsed = getIsoPPOBInquiry(serverResp)

	// Process PPOB Payment request
	case "810001":
		// Convert ISO message to JSON format
		jsonIso := getJsonPPOBPayment(msg)

		// Send JSON data to Biller
		serverResp := responsePPOBPayment(jsonIso)

		// Convert response from JSON data to ISO8583 format
		isoParsed = getIsoPPOBPayment(serverResp)

	// Process PPOB Status request
	case "380002":
		// Convert ISO message to JSON format
		jsonIso := getJsonPPOBStatus(msg)

		// Send JSON data to Biller
		serverResp := responsePPOBStatus(jsonIso)

		// Convert response from JSON data to ISO8583 format
		isoParsed = getIsoPPOBStatus(serverResp)

	// Process Topup Buy
	case "810002":
		// Convert ISO message to JSON format
		jsonIso := getJsonTopupBuy(msg)

		// Send JSON data to Biller
		serverResp := responseTopupBuy(jsonIso)

		// Convert response from JSON data to ISO8583 format
		isoParsed = getIsoTopupBuy(serverResp)

	// Process Topup Check
	case "380003":
		// Convert ISO message to JSON format
		jsonIso := getJsonTopupCheck(msg)

		// Send JSON data to Biller
		serverResp := responseTopupCheck(jsonIso)

		// Convert response from JSON data to ISO8583 format
		isoParsed = getIsoTopupCheck(serverResp)
	}

	isoMessage, _ := isoParsed.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))

	response.Header, _ = strconv.Atoi(isoHeader)
	response.MTI = isoParsed.Mti.String()
	response.Hex, _ = iso8583.BitMapArrayToHex(isoParsed.Bitmap)
	response.Message = isoMessage

	isoResponse = isoHeader + isoMessage
	log.Printf("\n\nResponse: \n\tHeader: %v\n\tMTI: %v\n\tHex: %v\n\tIso Message: %v\n\tFull Message: %v\n\n",
		response.Header,
		response.MTI,
		response.Hex,
		response.Message,
		isoResponse)

	// create file from response
	filename := "Response_to_" + isoParsed.Elements.GetElements()[3] + "@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	file := CreateFile("storage/response/"+filename, isoResponse)
	log.Println("File created: ", file)

	return isoResponse

}
