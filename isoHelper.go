package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/mofax/iso8583"
)

// Return ISO Message by converting data from map[int]string
func getIso(data map[int]string, mti string) (iso iso8583.IsoStruct) {
	log.Println("Converting to ISO8583...")

	isoStruct := iso8583.NewISOStruct("spec1987.yml", true)
	spec, _ := specFromFile("spec1987.yml")

	if isoStruct.Mti.String() != "" {
		log.Printf("Empty generates invalid MTI")
	}

	// Compare request data length and spec data length, add padding if different
	for field, data := range data {

		fieldSpec := spec.fields[field]

		// Check length for field with Length Type "fixed"
		if fieldSpec.LenType == "fixed" {
			lengthValidate, _ := iso8583.FixedLengthIntegerValidator(field, fieldSpec.MaxLen, data)

			if lengthValidate == false {
				if fieldSpec.ContentType == "n" {
					// Add padding for numeric field
					data = leftPad(data, fieldSpec.MaxLen, "0")
				} else {
					// Add padding for non-numeric field
					data = rightPad(data, fieldSpec.MaxLen, " ")
				}
			}
		}

		// Add data to isoStruct
		isoStruct.AddField(int64(field), data)
	}

	// Add MTI to isoStruct
	isoStruct.AddMTI(mti)

	// Logging isoStruct field and value
	printSortedDE(isoStruct)

	return isoStruct
}

// Return ISO message for PPOB Status JSON response
func getIsoPPOBStatus(jsonResponse PPOBStatusResponse) iso8583.IsoStruct {

	log.Println("Converting PPOB Status JSON Response to ISO8583")
	log.Printf("PPOB Status Response (JSON): %v\n", jsonResponse)

	// Assign data to map and add MTI
	struk := strings.Join(jsonResponse.Struk, ",")
	var response map[int]string
	if jsonResponse.Rc == "00" {
		response = map[int]string{
			4:   strconv.Itoa(jsonResponse.Tagihan),
			5:   strconv.Itoa(jsonResponse.Admin),
			6:   strconv.Itoa(jsonResponse.TotalTagihan),
			37:  jsonResponse.Reffid,
			39:  jsonResponse.Rc,
			43:  jsonResponse.Nama,
			48:  jsonResponse.TglLunas,
			62:  struk,
			120: jsonResponse.Msg,
			121: jsonResponse.Produk,
			122: jsonResponse.Nopel,
			123: jsonResponse.ReffNo,
			124: jsonResponse.Status,
		}
	} else {
		response = map[int]string{
			39:  jsonResponse.Rc,
			48:  jsonResponse.Restime,
			120: jsonResponse.Msg,
		}
	}
	mti := "0210"

	// Converting request map to isoStruct
	isoStruct := getIso(response, mti)

	// Adding PAN for PPOB Status Response
	isoStruct.AddField(3, "380002")
	isoMessage, _ := isoStruct.ToString()

	log.Println("Convert Success")
	log.Printf("PPOB Status Response (ISO8583): %s\n", isoMessage)
	return isoStruct

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
