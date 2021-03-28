package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mofax/iso8583"
)

// Return JSON for PPOB Inquiry ISO message request
func getJsonPPOBInquiry(parsedIso iso8583.IsoStruct) PPOBInquiryRequest {
	var response PPOBInquiryRequest

	log.Println("Converting PPOB Inquiry ISO8583 request to JSON")

	request, _ := parsedIso.ToString()
	log.Printf("PPOB Inquiry Request (ISO8583): %v\n", request)

	// Map ISO8583 format to JSON data
	emap := parsedIso.Elements.GetElements()
	response.TransactionID = strings.Trim(emap[48][0:25], " ")
	response.PartnerID = strings.Trim(emap[48][25:41], " ")
	response.ProductCode = strings.Trim(emap[48][41:57], " ")
	response.CustomerNo = strings.Trim(emap[48][57:82], " ")
	response.MerchantCode = strings.Trim(emap[48][82:107], " ")
	response.RequestTime = strings.Trim(emap[48][107:126], " ")
	response.Periode = strings.Trim(emap[48][126:], " ")

	// Create signature for new request
	signature := fmt.Sprintf("$inquiry$%v$%v$%v$%v$unand$",
		response.TransactionID, response.PartnerID, response.MerchantCode, response.RequestTime)
	log.Println("Signature:", signature)
	response.Signature = signatureSHA256(signature)
	log.Println("Signature encrypted: ", response.Signature)

	log.Println("Convert success")
	log.Printf("PPOB Inquiry Request (JSON): %+v\n", response)
	return response
}

// Return JSON for PPOB Payment ISO message request
func getJsonPPOBPayment(parsedIso iso8583.IsoStruct) PPOBPaymentRequest {
	var response PPOBPaymentRequest

	log.Println("Converting PPOB Payment ISO8583 request to JSON")

	request, _ := parsedIso.ToString()
	log.Printf("PPOB Payment Request (ISO8583): %v\n", request)

	// Map ISO8583 format to JSON data
	emap := parsedIso.Elements.GetElements()
	response.Amount, _ = strconv.Atoi(emap[4])
	response.ReffID = strings.Trim(emap[37], " ")
	response.TransactionID = strings.Trim(emap[48][0:25], " ")
	response.PartnerID = strings.Trim(emap[48][25:41], " ")
	response.ProductCode = strings.Trim(emap[48][41:57], " ")
	response.CustomerNo = strings.Trim(emap[48][57:82], " ")
	response.MerchantCode = strings.Trim(emap[48][82:107], " ")
	response.RequestTime = strings.Trim(emap[48][107:126], " ")

	// Create signature for new request
	signature := fmt.Sprintf("$payment$%v$%v$%v$%v$%v$unand$",
		response.TransactionID, response.PartnerID, response.ReffID, response.MerchantCode, response.RequestTime)
	log.Println("Signature: ", signature)
	response.Signature = signatureSHA256(signature)
	log.Println("Signature encrypted: ", response.Signature)

	log.Println("Convert success")
	log.Printf("PPOB Payment Request (JSON): %+v\n", response)
	return response
}

// Return JSON for Topup Buy ISO message request
func getJsonTopupBuy(parsedIso iso8583.IsoStruct) TopupBuyRequest {
	var response TopupBuyRequest

	log.Println("Converting Topup Buy ISO8583 request to JSON")

	request, _ := parsedIso.ToString()
	log.Printf("Topup Buy Request (ISO8583): %v\n", request)

	// Map ISO8583 format to JSON data
	emap := parsedIso.Elements.GetElements()
	response.TransactionID = strings.Trim(emap[48][0:25], " ")
	response.PartnerID = strings.Trim(emap[48][25:41], " ")
	response.ProductCode = strings.Trim(emap[48][41:57], " ")
	response.CustomerNo = strings.Trim(emap[48][57:82], " ")
	response.MerchantCode = strings.Trim(emap[48][82:107], " ")
	response.RequestTime = strings.Trim(emap[48][107:126], " ")

	// Create signature for new request
	signature := fmt.Sprintf("$buy$%v$%v$%v$%v$unand$",
		response.TransactionID, response.PartnerID, response.MerchantCode, response.RequestTime)
	log.Println("Signature: ", signature)
	response.Signature = signatureSHA256(signature)
	log.Println("Signature encrypted: ", response.Signature)

	log.Println("Convert success")
	log.Printf("Topup Buy Request (JSON): %+v\n", response)
	return response
}

// Return JSON for Topup Check ISO message request
func getJsonTopupCheck(parsedIso iso8583.IsoStruct) TopupCheckRequest {
	var response TopupCheckRequest

	log.Println("Converting Topup Check ISO8583 request to JSON")

	request, _ := parsedIso.ToString()
	log.Printf("Topup Check Request (ISO8583): %v\n", request)

	// Map ISO8583 format to JSON data
	emap := parsedIso.Elements.GetElements()
	response.TransactionID = strings.Trim(emap[48][0:25], " ")
	response.PartnerID = strings.Trim(emap[48][25:41], " ")
	response.ProductCode = strings.Trim(emap[48][41:57], " ")
	response.CustomerNo = strings.Trim(emap[48][57:82], " ")
	response.MerchantCode = strings.Trim(emap[48][82:107], " ")
	response.RequestTime = strings.Trim(emap[48][107:126], " ")

	// Create signature for new request
	signature := fmt.Sprintf("$check$%v$%v$%v$%v$unand$",
		response.TransactionID, response.PartnerID, response.MerchantCode, response.RequestTime)
	log.Println("Signature: ", signature)
	response.Signature = signatureSHA256(signature)
	log.Println("Signature encrypted: ", response.Signature)

	log.Println("Convert success")
	log.Printf("Topup Check Request (JSON): %+v\n", response)
	return response
}

// Return JSON for PPOB Status ISO message request
func getJsonPPOBStatus(parsedIso iso8583.IsoStruct) PPOBStatusRequest {
	var response PPOBStatusRequest

	log.Println("Converting PPOB Status ISO8583 request to JSON")

	request, _ := parsedIso.ToString()
	log.Printf("PPOB Status Request (ISO8583): %v\n", request)

	// Map ISO8583 format to JSON data
	emap := parsedIso.Elements.GetElements()
	response.Amount, _ = strconv.Atoi(emap[4])
	response.ReffID = strings.Trim(emap[37], " ")
	response.TransactionID = strings.Trim(emap[48][0:25], " ")
	response.PartnerID = strings.Trim(emap[48][25:41], " ")
	response.ProductCode = strings.Trim(emap[48][41:57], " ")
	response.CustomerNo = strings.Trim(emap[48][57:82], " ")
	response.MerchantCode = strings.Trim(emap[48][82:107], " ")
	response.RequestTime = strings.Trim(emap[48][107:126], " ")

	// Create signature for new request
	signature := fmt.Sprintf("$status$%v$%v$%v$%v$%v$unand$",
		response.TransactionID, response.PartnerID, response.ReffID, response.MerchantCode, response.RequestTime)
	log.Println("Signature: ", signature)
	response.Signature = signatureSHA256(signature)
	log.Println("Signature encrypted: ", response.Signature)

	log.Println("Convert success")
	log.Printf("PPOB Status Request (JSON): %+v\n", response)
	return response
}

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
