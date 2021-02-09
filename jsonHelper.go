package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mofax/iso8583"
)

// Any helper to process JSON data
// converter, formatter, etc

// Conver ISO message to JSON PPOBInquiry
func convJsonPPOBInquiry(parsedIso iso8583.IsoStruct) PPOBInquiryRequest {
	var response PPOBInquiryRequest

	log.Println("Converting ISO8583 to JSON")

	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.TransactionID = strings.Trim(emap[48][0:25], " ")
	response.PartnerID = strings.Trim(emap[48][25:41], " ")
	response.ProductCode = strings.Trim(emap[48][41:57], " ")
	response.CustomerNo = strings.Trim(emap[48][57:82], " ")
	response.MerchantCode = strings.Trim(emap[48][82:107], " ")
	response.RequestTime = strings.Trim(emap[48][107:126], " ")
	response.Periode = strings.Trim(emap[48][126:], " ")

	signature := fmt.Sprintf("$inquiry$%v$%v$%v$%v$unand$",
		response.TransactionID, response.PartnerID, response.MerchantCode, response.RequestTime)
	log.Println("Signature: ", signature)
	response.Signature = signatureSHA256(signature)
	log.Println("Signature encrypted: ", response.Signature)

	log.Printf("%+v\n", response)
	log.Println("Convert success")
	return response
}

// Conver ISO message to JSON PPOBPayment
func convJsonPPOBPayment(parsedIso iso8583.IsoStruct) PPOBPaymentRequest {
	var response PPOBPaymentRequest

	log.Println("Converting ISO8583 to JSON")

	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.Amount, _ = strconv.Atoi(emap[4])
	response.ReffID = strings.Trim(emap[37], " ")
	response.TransactionID = strings.Trim(emap[48][0:25], " ")
	response.PartnerID = strings.Trim(emap[48][25:41], " ")
	response.ProductCode = strings.Trim(emap[48][41:57], " ")
	response.CustomerNo = strings.Trim(emap[48][57:82], " ")
	response.MerchantCode = strings.Trim(emap[48][82:107], " ")
	response.RequestTime = strings.Trim(emap[48][107:126], " ")

	signature := fmt.Sprintf("$payment$%v$%v$%v$%v$%v$unand$",
		response.TransactionID, response.PartnerID, response.ReffID, response.MerchantCode, response.RequestTime)
	log.Println("Signature: ", signature)
	response.Signature = signatureSHA256(signature)
	log.Println("Signature encrypted: ", response.Signature)

	log.Printf("%+v\n", response)
	log.Println("Convert success")
	return response
}

// Conver ISO message to JSON TopupBuy
func convJsonTopupBuy(parsedIso iso8583.IsoStruct) TopupBuyRequest {
	var response TopupBuyRequest

	log.Println("Converting ISO8583 to JSON")

	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.TransactionID = strings.Trim(emap[48][0:25], " ")
	response.PartnerID = strings.Trim(emap[48][25:41], " ")
	response.ProductCode = strings.Trim(emap[48][41:57], " ")
	response.CustomerNo = strings.Trim(emap[48][57:82], " ")
	response.MerchantCode = strings.Trim(emap[48][82:107], " ")
	response.RequestTime = strings.Trim(emap[48][107:126], " ")

	signature := fmt.Sprintf("$buy$%v$%v$%v$%v$unand$",
		response.TransactionID, response.PartnerID, response.MerchantCode, response.RequestTime)
	log.Println("Signature: ", signature)
	response.Signature = signatureSHA256(signature)
	log.Println("Signature encrypted: ", response.Signature)

	log.Printf("%+v\n", response)
	log.Println("Convert success")
	return response
}

// Conver ISO message to JSON TopupCheck
func convJsonTopupCheck(parsedIso iso8583.IsoStruct) TopupCheckRequest {
	var response TopupCheckRequest

	log.Println("Converting ISO8583 to JSON")

	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.TransactionID = strings.Trim(emap[48][0:25], " ")
	response.PartnerID = strings.Trim(emap[48][25:41], " ")
	response.ProductCode = strings.Trim(emap[48][41:57], " ")
	response.CustomerNo = strings.Trim(emap[48][57:82], " ")
	response.MerchantCode = strings.Trim(emap[48][82:107], " ")
	response.RequestTime = strings.Trim(emap[48][107:126], " ")

	signature := fmt.Sprintf("$check$%v$%v$%v$%v$unand$",
		response.TransactionID, response.PartnerID, response.MerchantCode, response.RequestTime)
	log.Println("Signature: ", signature)
	response.Signature = signatureSHA256(signature)
	log.Println("Signature encrypted: ", response.Signature)

	log.Printf("%+v\n", response)
	log.Println("Convert success")
	return response
}

// Conver ISO message to JSON PPOBStatus
func convJsonPPOBStatus(parsedIso iso8583.IsoStruct) PPOBStatusRequest {
	var response PPOBStatusRequest

	log.Println("Converting ISO8583 to JSON")

	emap := parsedIso.Elements.GetElements()

	// Map ISO8583 format to JSON data
	response.Amount, _ = strconv.Atoi(emap[4])
	response.ReffID = strings.Trim(emap[37], " ")
	response.TransactionID = strings.Trim(emap[48][0:25], " ")
	response.PartnerID = strings.Trim(emap[48][25:41], " ")
	response.ProductCode = strings.Trim(emap[48][41:57], " ")
	response.CustomerNo = strings.Trim(emap[48][57:82], " ")
	response.MerchantCode = strings.Trim(emap[48][82:107], " ")
	response.RequestTime = strings.Trim(emap[48][107:126], " ")

	signature := fmt.Sprintf("$status$%v$%v$%v$%v$%v$unand$",
		response.TransactionID, response.PartnerID, response.ReffID, response.MerchantCode, response.RequestTime)
	log.Println("Signature: ", signature)
	response.Signature = signatureSHA256(signature)
	log.Println("Signature encrypted: ", response.Signature)

	log.Printf("%+v\n", response)
	log.Println("Convert success")
	return response
}

// Convert ISO message to JSON
func convertIsoToJson(parsedIso iso8583.IsoStruct) Transaction {
	var response Transaction

	log.Println("Converting ISO8583 to JSON")

	emap := parsedIso.Elements.GetElements()

	// Format padded CardAcceptor data
	if emap[43] != "" {
		cardAcceptorTerminalId := strings.TrimRight(emap[41], " ")
		cardAcceptorName := strings.TrimRight(emap[43][:25], " ")
		cardAcceptorCity := strings.TrimRight(emap[43][25:38], " ")
		cardAcceptorCountryCode := strings.TrimRight(emap[43][38:], " ")
		response.CardAcceptorData.CardAcceptorTerminalId = cardAcceptorTerminalId
		response.CardAcceptorData.CardAcceptorName = cardAcceptorName
		response.CardAcceptorData.CardAcceptorCity = cardAcceptorCity
		response.CardAcceptorData.CardAcceptorCountryCode = cardAcceptorCountryCode
	}

	// Map ISO8583 format to JSON data
	response.Pan = emap[2]
	response.ProcessingCode = emap[3]
	response.TotalAmount, _ = strconv.Atoi(emap[4])
	response.SettlementAmount = emap[5]
	response.CardholderBillingAmount = emap[6]
	response.TransmissionDateTime = emap[7]
	response.SettlementConversionRate = emap[9]
	response.CardHolderBillingConvRate = emap[10]
	response.Stan = emap[11]
	response.LocalTransactionTime = emap[12]
	response.LocalTransactionDate = emap[13]
	response.CaptureDate = emap[17]
	response.CategoryCode = emap[18]
	response.PointOfServiceEntryMode = emap[22]
	response.Refnum = emap[37]
	response.AdditionalData = emap[48]
	response.Currency = emap[49]
	response.SettlementCurrencyCode = emap[50]
	response.CardHolderBillingCurrencyCode = emap[51]
	response.AdditionalDataNational = emap[57]

	log.Println("Convert success")
	return response
}
