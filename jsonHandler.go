package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Handle all JSON Client request

// Send request to mock server in JSON format
// Get response from mock server in JSON format
func responseJson(jsonIso Transaction) PaymentResponse {
	var response PaymentResponse

	// Initiate request body
	requestBody, err := json.Marshal(jsonIso)
	if err != nil {
		log.Fatalf("Preparing body request failed. Error: %v\n", err)
	}

	// Client setup for custom http request
	client := &http.Client{}

	log.Printf("Request to https://tiruan.herokuapp.com/biller\n")

	// Request to mock server
	req, err := http.NewRequest("GET", "https://tiruan.herokuapp.com/biller", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalf("Failed to sent request to https://tiruan.herokuapp.com/biller. Error: %v\n", err)
	}

	// Check response from mock server
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to get response from https://tiruan.herokuapp.com/biller. Error: %v\n", err)
	}

	defer resp.Body.Close()

	log.Printf("Response from https://tiruan.herokuapp.com/biller\n")

	// Read response from mock server
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	// Get transactionData from mock server response

	return response
}

func mockGo(jsonIso GoRoutineReq) GoRoutineRes {
	var response GoRoutineRes

	req := jsonIso.Data
	response.Response = "Response: " + req

	return response
}

// Get response from mock server in JSON format
func responseJsonPPOBInquiry(jsonIso PPOBInquiryRequest) PPOBInquiryResponse {
	var response PPOBInquiryResponse

	// Initiate request body
	//requestBody, err := json.Marshal(jsonIso)
	//if err != nil {
	//	log.Fatalf("Preparing body request failed. Error: %v\n", err)
	//}

	// Client setup for custom http request
	client := &http.Client{}
	var baseURL = "https://chipsakti-mock.herokuapp.com"

	// Set data to be encoded
	var param = url.Values{}
	param.Set("transaction_id", jsonIso.TransactionID)
	param.Set("partner_id", jsonIso.PartnerID)
	param.Set("product_code", jsonIso.ProductCode)
	param.Set("customer_no", jsonIso.CustomerNo)
	param.Set("periode", jsonIso.Periode)
	param.Set("merchant_code", jsonIso.MerchantCode)
	param.Set("request_time", jsonIso.RequestTime)
	param.Set("signature", jsonIso.Signature)
	sig := param.Get("signature")
	log.Println("Sent signature : ", sig)
	reqTime := param.Get("request_time")
	log.Println("Sent request time : ", reqTime)
	var payload = bytes.NewBufferString(param.Encode())

	log.Printf("Request to https://chipsakti-mock.herokuapp.com/inquiry\n")

	// Request to mock server
	req, err := http.NewRequest("POST", baseURL+"/inquiry", payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Fatalf("Failed to sent request to https://chipsakti-mock.herokuapp.com/inquiry. Error: %v\n", err)
	}

	// Check response from mock server
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to get response from https://chipsakti-mock.herokuapp.com/inquiry. Error: %v\n", err)
	}

	defer resp.Body.Close()

	log.Printf("Response from https://chipsakti-mock.herokuapp.com/inquiry\n")

	// Read response from mock server
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	// Get PPOBInquiryResponse from mock server response

	return response
}

func responsePPOBPayment(jsonIso PPOBPaymentRequest) PPOBPaymentResponse {
	var response PPOBPaymentResponse

	// Initiate request body
	//requestBody, err := json.Marshal(jsonIso)
	//if err != nil {
	//	log.Fatalf("Preparing body request failed. Error: %v\n", err)
	//}

	// Client setup for custom http request
	client := &http.Client{}
	var baseURL = "https://chipsakti-mock.herokuapp.com"
	amount := strconv.Itoa(jsonIso.Amount)

	// Set data to be encoded
	var param = url.Values{}
	param.Set("transaction_id", jsonIso.TransactionID)
	param.Set("partner_id", jsonIso.PartnerID)
	param.Set("product_code", jsonIso.ProductCode)
	param.Set("customer_no", jsonIso.CustomerNo)
	param.Set("reff_id", jsonIso.ReffID)
	param.Set("amount", amount)
	param.Set("merchant_code", jsonIso.MerchantCode)
	param.Set("request_time", jsonIso.RequestTime)
	param.Set("signature", jsonIso.Signature)
	sig := param.Get("signature")
	log.Println("Sent signature : ", sig)
	reqTime := param.Get("request_time")
	log.Println("Sent request time : ", reqTime)
	var payload = bytes.NewBufferString(param.Encode())

	log.Printf("Request to https://chipsakti-mock.herokuapp.com/payment\n")

	// Request to mock server
	req, err := http.NewRequest("POST", baseURL+"/payment", payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Fatalf("Failed to sent request to https://chipsakti-mock.herokuapp.com/payment. Error: %v\n", err)
	}

	// Check response from mock server
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to get response from https://chipsakti-mock.herokuapp.com/payment. Error: %v\n", err)
	}

	defer resp.Body.Close()

	log.Printf("Response from https://chipsakti-mock.herokuapp.com/payment\n")

	// Read response from mock server
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	// Get PPOBInquiryResponse from mock server response

	return response
}

func responsePPOBStatus(jsonIso PPOBStatusRequest) PPOBStatusResponse {
	var response PPOBStatusResponse

	// Initiate request body
	//requestBody, err := json.Marshal(jsonIso)
	//if err != nil {
	//	log.Fatalf("Preparing body request failed. Error: %v\n", err)
	//}

	// Client setup for custom http request
	client := &http.Client{}
	var baseURL = "https://chipsakti-mock.herokuapp.com"
	amount := strconv.Itoa(jsonIso.Amount)

	// Set data to be encoded
	var param = url.Values{}
	param.Set("transaction_id", jsonIso.TransactionID)
	param.Set("partner_id", jsonIso.PartnerID)
	param.Set("product_code", jsonIso.ProductCode)
	param.Set("customer_no", jsonIso.CustomerNo)
	param.Set("reff_id", jsonIso.ReffID)
	param.Set("amount", amount)
	param.Set("merchant_code", jsonIso.MerchantCode)
	param.Set("request_time", jsonIso.RequestTime)
	param.Set("signature", jsonIso.Signature)
	sig := param.Get("signature")
	log.Println("Sent signature : ", sig)
	reqTime := param.Get("request_time")
	log.Println("Sent request time : ", reqTime)
	var payload = bytes.NewBufferString(param.Encode())

	log.Printf("Request to https://chipsakti-mock.herokuapp.com/status\n")

	// Request to mock server
	req, err := http.NewRequest("POST", baseURL+"/status", payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Fatalf("Failed to sent request to https://chipsakti-mock.herokuapp.com/status. Error: %v\n", err)
	}

	// Check response from mock server
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to get response from https://chipsakti-mock.herokuapp.com/status. Error: %v\n", err)
	}

	defer resp.Body.Close()

	log.Printf("Response from https://chipsakti-mock.herokuapp.com/status\n")

	// Read response from mock server
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	// Get PPOBInquiryResponse from mock server response

	return response
}

func responseTopupBuy(jsonIso TopupBuyRequest) TopupBuyResponse {
	var response TopupBuyResponse

	// Initiate request body
	//requestBody, err := json.Marshal(jsonIso)
	//if err != nil {
	//	log.Fatalf("Preparing body request failed. Error: %v\n", err)
	//}

	// Client setup for custom http request
	client := &http.Client{}
	var baseURL = "https://chipsakti-mock.herokuapp.com"

	// Set data to be encoded
	var param = url.Values{}
	param.Set("transaction_id", jsonIso.TransactionID)
	param.Set("partner_id", jsonIso.PartnerID)
	param.Set("product_code", jsonIso.ProductCode)
	param.Set("customer_no", jsonIso.CustomerNo)
	param.Set("merchant_code", jsonIso.MerchantCode)
	param.Set("request_time", jsonIso.RequestTime)
	param.Set("signature", jsonIso.Signature)
	sig := param.Get("signature")
	log.Println("Sent signature : ", sig)
	reqTime := param.Get("request_time")
	log.Println("Sent request time : ", reqTime)
	var payload = bytes.NewBufferString(param.Encode())

	log.Printf("Request to https://chipsakti-mock.herokuapp.com/buy\n")

	// Request to mock server
	req, err := http.NewRequest("POST", baseURL+"/buy", payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Fatalf("Failed to sent request to https://chipsakti-mock.herokuapp.com/buy. Error: %v\n", err)
	}

	// Check response from mock server
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to get response from https://chipsakti-mock.herokuapp.com/buy. Error: %v\n", err)
	}

	defer resp.Body.Close()

	log.Printf("Response from https://chipsakti-mock.herokuapp.com/buy\n")

	// Read response from mock server
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	// Get PPOBInquiryResponse from mock server response

	return response
}

func responseTopupCheck(jsonIso TopupCheckRequest) TopupCheckResponse {
	var response TopupCheckResponse

	// Initiate request body
	//requestBody, err := json.Marshal(jsonIso)
	//if err != nil {
	//	log.Fatalf("Preparing body request failed. Error: %v\n", err)
	//}

	// Client setup for custom http request
	client := &http.Client{}
	var baseURL = "https://chipsakti-mock.herokuapp.com"

	// Set data to be encoded
	var param = url.Values{}
	param.Set("transaction_id", jsonIso.TransactionID)
	param.Set("partner_id", jsonIso.PartnerID)
	param.Set("product_code", jsonIso.ProductCode)
	param.Set("customer_no", jsonIso.CustomerNo)
	param.Set("merchant_code", jsonIso.MerchantCode)
	param.Set("request_time", jsonIso.RequestTime)
	param.Set("signature", jsonIso.Signature)
	sig := param.Get("signature")
	log.Println("Sent signature : ", sig)
	reqTime := param.Get("request_time")
	log.Println("Sent request time : ", reqTime)
	var payload = bytes.NewBufferString(param.Encode())

	log.Printf("Request to https://chipsakti-mock.herokuapp.com/check\n")

	// Request to mock server
	req, err := http.NewRequest("POST", baseURL+"/check", payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Fatalf("Failed to sent request to https://chipsakti-mock.herokuapp.com/check. Error: %v\n", err)
	}

	// Check response from mock server
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to get response from https://chipsakti-mock.herokuapp.com/check. Error: %v\n", err)
	}

	defer resp.Body.Close()

	log.Printf("Response from https://chipsakti-mock.herokuapp.com/check\n")

	// Read response from mock server
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	// Get PPOBInquiryResponse from mock server response

	return response
}
