package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mofax/iso8583"
	"github.com/rivo/uniseg"
)

// Handle all ISO Client request

// Process ISO message in body request
func sendIso(writer http.ResponseWriter, request *http.Request) {

	var response Response
	var iso Iso8583

	// Read body request
	reqBody, _ := ioutil.ReadAll(request.Body)
	req := string(reqBody)
	log.Printf("ISO Message: %v\n", req)

	// Produce event
	err := doProducer(broker, topic1, req)

	if err != nil {
		errDesc := fmt.Sprintf("Failed sent to Kafka\nError: %v", err)
		log.Println(err)
		response.ResponseCode, response.ResponseDescription = 500, errDesc
		jsonFormatter(writer, response, 500)
	} else {
		// Read response
		msg, err := consumeResponse(broker, group, []string{topic2})
		if err != nil {
			errDesc := fmt.Sprintf("Failed to get response from Kafka\nError: %v", err)
			log.Println(err)
			response.ResponseCode, response.ResponseDescription = 500, errDesc
			jsonFormatter(writer, response, 500)
		} else {

			// Return empty response
			if msg == "" {
				errDesc := "Got empty response"
				log.Println(errDesc)
				response.ResponseCode, response.ResponseDescription = 500, errDesc
				jsonFormatter(writer, response, 500)
			} else {

				// Parse response string to ISO8583 data
				header := msg[0:4]
				data := msg[4:]

				isoStruct := iso8583.NewISOStruct("spec1987.yml", false)

				isoParsed, err := isoStruct.Parse(data)
				if err != nil {
					log.Printf("Error parsing iso message\nError: %v", err)
				}

				iso.Header, _ = strconv.Atoi(header)
				iso.MTI = isoParsed.Mti.String()
				iso.Hex, _ = iso8583.BitMapArrayToHex(isoParsed.Bitmap)

				iso.Message, err = isoParsed.ToString()
				if err != nil {
					log.Printf("Iso Parsed failed convert to string.\nError: %v", err)
				}

				//event := header + iso.Message

				iso.ResponseStatus.ResponseCode, iso.ResponseStatus.ResponseDescription = 200, "Success"
				jsonFormatter(writer, iso, 200)

			}

		}
	}

}

// Get response from mock server in ISO Format
func responseIso(message string) {

	var response Iso8583
	data := message[4:]

	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)

	msg, err := isoStruct.Parse(data)
	if err != nil {
		log.Println(err)
	}

	var isoParsed iso8583.IsoStruct

	// Check processing code
	pcode := msg.Elements.GetElements()[3]
	if pcode == "380001" {
		// Convert ISO message to JSON format
		jsonIso := convJsonPPOBInquiry(msg)

		// Send JSON data to mock server
		serverResp := responsePPOBInquiry(jsonIso)

		// Convert response from JSON data to ISO8583 format
		isoParsed = convIsoPPOBInquiry(serverResp)

		isoParsed.AddField(3, "380001")
	} else if pcode == "810001" {
		// Convert ISO message to JSON format
		jsonIso := convJsonPPOBPayment(msg)

		// Send JSON data to mock server
		serverResp := responsePPOBPayment(jsonIso)

		// Convert response from JSON data to ISO8583 format
		isoParsed = convIsoPPOBPayment(serverResp)

		isoParsed.AddField(3, "810001")
	} else if pcode == "380002" {
		// Convert ISO message to JSON format
		jsonIso := convJsonPPOBStatus(msg)

		// Send JSON data to mock server
		serverResp := responsePPOBStatus(jsonIso)

		// Convert response from JSON data to ISO8583 format
		isoParsed = convIsoPPOBStatus(serverResp)

		isoParsed.AddField(3, "380002")
	} else if pcode == "810002" {
		// Convert ISO message to JSON format
		jsonIso := convJsonTopupBuy(msg)

		// Send JSON data to mock server
		serverResp := responseTopupBuy(jsonIso)

		// Convert response from JSON data to ISO8583 format
		isoParsed = convIsoTopupBuy(serverResp)

		isoParsed.AddField(3, "810002")
	} else if pcode == "380003" {
		// Convert ISO message to JSON format
		jsonIso := convJsonTopupCheck(msg)

		// Send JSON data to mock server
		serverResp := responseTopupCheck(jsonIso)

		// Convert response from JSON data to ISO8583 format
		isoParsed = convIsoTopupCheck(serverResp)

		isoParsed.AddField(3, "380003")
	}

	// Change MTI response
	isoParsed.AddMTI("0210")

	isoMessage, _ := isoParsed.ToString()
	isoHeader := fmt.Sprintf("%04d", uniseg.GraphemeClusterCount(isoMessage))

	response.Header, _ = strconv.Atoi(isoHeader)
	response.MTI = isoParsed.Mti.String()
	response.Hex, _ = iso8583.BitMapArrayToHex(isoParsed.Bitmap)
	response.Message = isoMessage

	event := isoHeader + isoMessage
	log.Printf("\n\nResponse: \n\tHeader: %v\n\tMTI: %v\n\tHex: %v\n\tIso Message: %v\n\tFull Message: %v\n\n",
		response.Header,
		response.MTI,
		response.Hex,
		response.Message,
		event)

	// create file from response
	filename := "Response_to_" + isoParsed.Elements.GetElements()[3] + "@" + fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05"))
	file := CreateFile("storage/response/"+filename, event)
	log.Println("File created: ", file)

	// Produce event
	err = doProducer(broker, topic4, event)
	if err != nil {
		log.Printf("Error producing message %v\n", message)
		log.Println(err)
	}

}
