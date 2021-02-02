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

func responsePPOBStatus(jsonIso PPOBStatusRequest) PPOBStatusResponse {
	var response PPOBStatusResponse

	response.Produk = "INDOVISION"
	response.Nopel = "301124297"
	response.Nama = "YUNUS."
	response.Tagihan = 351456
	response.Admin = 0
	response.TotalTagihan = 351456
	response.TglLunas = "2018-07-02 15:54:08"
	response.ReffNo = "200557BN"
	response.Struk = "This is struk array"
	response.Msg = "Approve"
	response.Rc = "00"
	response.Status = "Payment Successful"
	response.Reffid = "11119"

	return response
}

func responseTopupBuy(jsonIso TopupBuyRequest) TopupBuyResponse {
	var response TopupBuyResponse

	response.Rc = "00"
	response.Msg = "Pembelian TSEL5 0812344321 pada 03 Jul 2018, 15:03 BERHASIL.SN=1530604784. Harga Rp. 5800;"
	response.SN = "1530604784"
	response.Price = "5800"
	response.Restime = "2018-05-15 15:10:05"

	return response
}

func responseTopupCheck(jsonIso TopupCheckRequest) TopupCheckResponse {
	var response TopupCheckResponse

	response.Rc = "00"
	response.Msg = "Pembelian TSEL5 0818337744 pada 03 Jul 2018, 15:03 BERHASIL. SN=1530604784. Harga Rp. 5800;"
	response.SN = "1530604784"
	response.Price = "5800"
	response.Restime = "2018-05-15 15:10:05"

	return response
}
