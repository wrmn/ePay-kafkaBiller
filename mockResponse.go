package main

// Send request to mock server in JSON format
// Get response from mock server in JSON format
func responsePPOBInquiry(jsonIso PPOBInquiryRequest) PPOBInquiryResponse {
	var response PPOBInquiryResponse

	response.Tagihan = 8700000
	response.Admin = 3300
	response.TotalTagihan = 873300
	response.Reffid = "10226"
	response.Rc = "00"
	response.Nama = "ANDIK SUNARKO"
	response.Restime = "2018-07-01 12:32:20"
	response.Data = "PPOBInquiryResponse"
	response.Msg = "Approve"
	response.Produk = "WOM"
	response.Nopel = "801100022085"

	return response
}

func responsePPOBPayment(jsonIso PPOBPaymentRequest) PPOBPaymentResponse {
	var response PPOBPaymentResponse

	response.Produk = "WOM"
	response.Nopel = "801100022085"
	response.Nama = "ANDIK SUNARKO"
	response.Tagihan = 870000
	response.Admin = 3300
	response.TotalTagihan = 873300
	response.TglLunas = "2018-06-25 15:21:56"
	response.ReffNo = "20160906029000023990"
	response.Struk = "This is struk array"
	response.Msg = "Approve"
	response.Rc = "00"
	response.Reffid = "10228"

	return response
}
