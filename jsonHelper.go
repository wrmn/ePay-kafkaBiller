package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/mofax/iso8583"
)

func getEpayRinstis(parsedIso iso8583.IsoStruct) (response rintisRequest) {

	log.Println("Converting PPOB Status ISO8583 request to JSON")

	request, _ := parsedIso.ToString()
	log.Printf("PPOB Status Request (ISO8583): %v\n", request)

	// Map ISO8583 format to JSON data
	emap := parsedIso.Elements.GetElements()
	response.Pan = strings.Trim(emap[2], " ")
	response.ProcessingCode = strings.Trim(emap[3], " ")
	response.TotalAmount, _ = strconv.Atoi(emap[4])
	response.TransmissionDateTime = strings.Trim(emap[7], " ")
	response.Stan = strings.Trim(emap[11], " ")
	response.LocalTransactionTime = strings.Trim(emap[12], " ")
	response.LocalTransactionDate = strings.Trim(emap[13], " ")
	response.CaptureDate = strings.Trim(emap[17], " ")
	response.AcquirerID = strings.Trim(emap[32], " ")
	response.Track2Data = strings.Trim(emap[35], " ")
	response.Refnum = strings.Trim(emap[37], " ")
	response.TerminalID = strings.Trim(emap[41], " ")
	response.CardAcceptorData = strings.Trim(emap[43], " ")
	response.AdditionalData = strings.Trim(emap[48], " ")
	response.Currency = strings.Trim(emap[49], " ")
	response.PIN = strings.Trim(emap[52], " ")
	response.TerminalData = strings.Trim(emap[60], " ")
	response.AccountTo = strings.Trim(emap[103], " ")
	response.TokenData = strings.Trim(emap[126], " ")

	log.Println("Convert success")
	log.Printf("PPOB Status Request (JSON): %+v\n", response)

	return response
}
