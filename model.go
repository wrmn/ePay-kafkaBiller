package main

type CardAcceptorData struct {
	CardAcceptorTerminalId  string `json:"cardAcceptorTerminalID"`
	CardAcceptorName        string `json:"cardAcceptorName"`
	CardAcceptorCity        string `json:"cardAcceptorCity"`
	CardAcceptorCountryCode string `json:"cardAcceptorCountryCode"`
}

type Transaction struct {
	Pan                           string           `json:"pan"`
	ProcessingCode                string           `json:"processingCode"`
	TotalAmount                   int              `json:"totalAmount"`
	AcquirerId                    string           `json:"acquirerId"`
	IssuerId                      string           `json:"issuerId"`
	TransmissionDateTime          string           `json:"transmissionDateTime"`
	LocalTransactionTime          string           `json:"localTransactionTime"`
	LocalTransactionDate          string           `json:"localTransactionDate"`
	CaptureDate                   string           `json:"captureDate"`
	AdditionalData                string           `json:"additionalData"`
	Stan                          string           `json:"stan"`
	Refnum                        string           `json:"refnum"`
	Currency                      string           `json:"currency"`
	TerminalId                    string           `json:"terminalId"`
	AccountFrom                   string           `json:"accountFrom"`
	AccountTo                     string           `json:"accountTo"`
	CategoryCode                  string           `json:"categoryCode"`
	SettlementAmount              string           `json:"settlementAmount"`
	CardholderBillingAmount       string           `json:"cardholderBillingAmount"`
	SettlementConversionRate      string           `json:"settlementConversionRate"`
	CardHolderBillingConvRate     string           `json:"cardHolderBillingConvRate"`
	PointOfServiceEntryMode       string           `json:"pointOfServiceEntryMode"`
	CardAcceptorData              CardAcceptorData `json:"cardAcceptorData"`
	SettlementCurrencyCode        string           `json:"settlementCurrencyCode"`
	CardHolderBillingCurrencyCode string           `json:"cardHolderBillingCurrencyCode"`
	AdditionalDataNational        string           `json:"additionalDataNational"`
}

type Response struct {
	ResponseCode        int    `json:"responseCode"`
	ReasonCode          int    `json:"reasonCode"`
	ResponseDescription string `json:"responseDescription"`
}

type PaymentResponse struct {
	TransactionData Transaction `json:"transactionData"`
	ResponseStatus  Response    `json:"responseStatus"`
}

type Iso8583 struct {
	Header         int      `json:"header"`
	MTI            string   `json:"mti"`
	Hex            string   `json:"hex"`
	Message        string   `json:"message"`
	ResponseStatus Response `json:"responseStatus"`
}

// Process ISO8583 Spec file
type fieldDescription struct {
	ContentType string `yaml:"ContentType"`
	MaxLen      int    `yaml:"MaxLen"`
	MinLen      int    `yaml:"MinLen"`
	LenType     string `yaml:"LenType"`
	Label       string `yaml:"Label"`
}

type PPOBInquiryRequest struct {
	TransactionID string `json:"transaction_id"`
	PartnerID     string `json:"partner_id"`
	ProductCode   string `json:"product_code"`
	CustomerNo    string `json:"customer_no"`
	Periode       string `json:"periode"`
	MerchantCode  string `json:"merchant_code"`
	RequestTime   string `json:"request_time"`
	Signature     string `json:"signature"`
}

type PPOBInquiryResponse struct {
	Rc           string `json:"rc"`
	Msg          string `json:"msg"`
	Produk       string `json:"produk"`
	Nopel        string `json:"nopel"`
	Nama         string `json:"nama"`
	Tagihan      int    `json:"tagihan"`
	Admin        int    `json:"admin"`
	TotalTagihan int    `json:"total_tagihan"`
	Reffid       string `json:"reffid"`
	Data         string `json:"data"`
	Restime      string `json:"restime"`
}

type PPOBPaymentRequest struct {
	TransactionID string `json:"transaction_id"`
	PartnerID     string `json:"partner_id"`
	ProductCode   string `json:"product_code"`
	CustomerNo    string `json:"customer_no"`
	MerchantCode  string `json:"merchant_code"`
	ReffID        string `json:"reff_id"`
	Amount        int    `json:"amount"`
	RequestTime   string `json:"request_time"`
	Signature     string `json:"signature"`
}

type PPOBPaymentResponse struct {
	Rc           string `json:"rc"`
	Msg          string `json:"msg"`
	Produk       string `json:"produk"`
	Nopel        string `json:"nopel"`
	Nama         string `json:"nama"`
	Tagihan      int    `json:"tagihan"`
	Admin        int    `json:"admin"`
	TotalTagihan int    `json:"total_tagihan"`
	Reffid       string `json:"reffid"`
	TglLunas     string `json:"tgl_lunas"`
	Struk        string `json:"struk"`
	ReffNo       string `json:"Reff_no"`
}
