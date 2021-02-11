package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/mofax/iso8583"
)

// Any helper to process ISO data
// converter, formatter, etc

// Convert JSON data to ISO8583 format
func convIsoPPOBInquiry(data PPOBInquiryResponse) iso8583.IsoStruct {
	log.Println("Converting JSON to ISO8583")

	trans := map[int64]string{
		4:   strconv.Itoa(data.Tagihan),
		5:   strconv.Itoa(data.Admin),
		6:   strconv.Itoa(data.TotalTagihan),
		37:  data.Reffid,
		39:  data.Rc,
		43:  data.Nama,
		48:  data.Restime,
		62:  data.Data,
		120: data.Msg,
		121: data.Produk,
		122: data.Nopel,
	}

	one := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if one.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	for field, data := range trans {

		fieldSpec := spec.fields[int(field)]

		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(int(field), fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		one.AddField(field, data)

	}

	printSortedDE(one)
	log.Println("Convert Success")
	return one

}

func convIsoPPOBPayment(data PPOBPaymentResponse) iso8583.IsoStruct {
	log.Println("Converting JSON to ISO8583")

	trans := map[int64]string{
		4:   strconv.Itoa(data.Tagihan),
		5:   strconv.Itoa(data.Admin),
		6:   strconv.Itoa(data.TotalTagihan),
		37:  data.Reffid,
		39:  data.Rc,
		43:  data.Nama,
		48:  data.TglLunas,
		62:  data.Struk,
		120: data.Msg,
		121: data.Produk,
		122: data.Nopel,
		123: data.ReffNo,
	}

	one := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if one.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	for field, data := range trans {

		fieldSpec := spec.fields[int(field)]

		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(int(field), fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		one.AddField(field, data)

	}

	printSortedDE(one)
	log.Println("Convert Success")
	return one

}

func convIsoPPOBStatus(data PPOBStatusResponse) iso8583.IsoStruct {
	log.Println("Converting JSON to ISO8583")

	trans := map[int64]string{
		4:   strconv.Itoa(data.Tagihan),
		5:   strconv.Itoa(data.Admin),
		6:   strconv.Itoa(data.TotalTagihan),
		37:  data.Reffid,
		39:  data.Rc,
		43:  data.Nama,
		48:  data.TglLunas,
		62:  data.Struk,
		120: data.Msg,
		121: data.Produk,
		122: data.Nopel,
		123: data.ReffNo,
		124: data.Status,
	}

	one := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if one.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	for field, data := range trans {

		fieldSpec := spec.fields[int(field)]

		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(int(field), fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		one.AddField(field, data)

	}

	printSortedDE(one)
	log.Println("Convert Success")
	return one

}

func convIsoTopupBuy(data TopupBuyResponse) iso8583.IsoStruct {
	log.Println("Converting JSON to ISO8583")

	trans := map[int64]string{
		39:  data.Rc,
		48:  data.Restime,
		120: data.Msg,
		121: data.SN,
		122: data.Price,
	}

	one := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if one.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	for field, data := range trans {

		fieldSpec := spec.fields[int(field)]

		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(int(field), fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		one.AddField(field, data)

	}

	printSortedDE(one)
	log.Println("Convert Success")
	return one

}

func convIsoTopupCheck(data TopupCheckResponse) iso8583.IsoStruct {
	log.Println("Converting JSON to ISO8583")

	trans := map[int64]string{
		39:  data.Rc,
		48:  data.Restime,
		120: data.Msg,
		121: data.SN,
		122: data.Price,
	}

	one := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if one.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	for field, data := range trans {

		fieldSpec := spec.fields[int(field)]

		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(int(field), fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		one.AddField(field, data)

	}

	printSortedDE(one)
	log.Println("Convert Success")
	return one

}

func convertJsonToIso(data PaymentResponse) iso8583.IsoStruct {

	log.Println("Converting JSON to ISO8583")

	cardAcceptorTerminalId := data.TransactionData.CardAcceptorData.CardAcceptorTerminalId
	cardAcceptorName := data.TransactionData.CardAcceptorData.CardAcceptorName
	cardAcceptorCity := data.TransactionData.CardAcceptorData.CardAcceptorCity
	cardAcceptorCountryCode := data.TransactionData.CardAcceptorData.CardAcceptorCountryCode
	responseStatus := convResp(data.ResponseStatus)

	if len(data.TransactionData.CardAcceptorData.CardAcceptorTerminalId) < 16 {
		cardAcceptorTerminalId = rightPad(data.TransactionData.CardAcceptorData.CardAcceptorTerminalId, 16, " ")
	}
	if len(data.TransactionData.CardAcceptorData.CardAcceptorName) < 25 {
		cardAcceptorName = rightPad(data.TransactionData.CardAcceptorData.CardAcceptorName, 25, " ")
	}
	if len(data.TransactionData.CardAcceptorData.CardAcceptorCity) < 13 {
		cardAcceptorCity = rightPad(data.TransactionData.CardAcceptorData.CardAcceptorCity, 13, " ")
	}
	if len(data.TransactionData.CardAcceptorData.CardAcceptorCountryCode) < 2 {
		cardAcceptorCountryCode = rightPad(data.TransactionData.CardAcceptorData.CardAcceptorCountryCode, 2, " ")
	}
	cardAcceptor := cardAcceptorName + cardAcceptorCity + cardAcceptorCountryCode

	trans := map[int64]string{
		2:  data.TransactionData.Pan,
		3:  data.TransactionData.ProcessingCode,
		4:  strconv.Itoa(data.TransactionData.TotalAmount),
		5:  data.TransactionData.SettlementAmount,
		6:  data.TransactionData.CardholderBillingAmount,
		7:  data.TransactionData.TransmissionDateTime,
		9:  data.TransactionData.SettlementConversionRate,
		10: data.TransactionData.CardHolderBillingConvRate,
		11: data.TransactionData.Stan,
		12: data.TransactionData.LocalTransactionTime,
		13: data.TransactionData.LocalTransactionDate,
		17: data.TransactionData.CaptureDate,
		18: data.TransactionData.CategoryCode,
		22: data.TransactionData.PointOfServiceEntryMode,
		37: data.TransactionData.Refnum,
		39: responseStatus,
		41: cardAcceptorTerminalId,
		43: cardAcceptor,
		48: data.TransactionData.AdditionalData,
		49: data.TransactionData.Currency,
		50: data.TransactionData.SettlementCurrencyCode,
		51: data.TransactionData.CardHolderBillingCurrencyCode,
		57: data.TransactionData.AdditionalDataNational,
	}

	one := iso8583.NewISOStruct("spec1987.yml", false)
	spec, _ := specFromFile("spec1987.yml")

	if one.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	for field, data := range trans {

		fieldSpec := spec.fields[int(field)]

		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(int(field), fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		one.AddField(field, data)

	}

	printSortedDE(one)
	log.Println("Convert Success")
	return one
}

// Log sorted converted ISO Message
func printSortedDE(parsedMessage iso8583.IsoStruct) {
	dataElement := parsedMessage.Elements.GetElements()
	int64toSort := make([]int, 0, len(dataElement))
	for key := range dataElement {
		int64toSort = append(int64toSort, int(key))
	}
	sort.Ints(int64toSort)
	for _, key := range int64toSort {
		log.Printf("[%v] : %v\n", int64(key), dataElement[int64(key)])
	}
}

// Spec contains a structured description of an iso8583 spec
// properly defined by a spec file
type Spec struct {
	fields map[int]fieldDescription
}

// readFromFile reads a yaml specfile and loads
// and iso8583 spec from it
func (s *Spec) readFromFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	yaml.Unmarshal(content, &s.fields) // expecting content to be valid yaml
	return nil
}

// SpecFromFile returns a brand new empty spec
func specFromFile(filename string) (Spec, error) {
	s := Spec{}
	err := s.readFromFile(filename)
	if err != nil {
		return s, err
	}
	return s, nil
}

// Add pad on left of data,
// Used to format number by adding "0" in front of number data
func leftPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return padding + s
}

// Add pad on right of data,
// Used to format string by adding " " at the end of string data
func rightPad(s string, length int, pad string) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(pad, length-len(s))
	return s + padding
}

func convResp(resp Response) string {
	response := fmt.Sprintf("%d%d%s", resp.ResponseCode, resp.ReasonCode, resp.ResponseDescription)
	res := len(response)
	return fmt.Sprintf("%d%s", res, response)
}
