package main221

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type soapRQ struct {
	XMLName   xml.Name `xml:"Envelope"`
	XMLNsSoap string   `xml:"xmlns,attr"`
	Body      soapBody `xml:"Body"`
}

type soapBody struct {
	ZFMWSPRECIOSTOCK ZFMWSPRECIOSTOCK `xml:"ZFM_WS_PRECIOSTOCK"`
}

type ZFMWSPRECIOSTOCK struct {
	INPUT    params `xml:"INPUT"`
	XMLNsURN string `xml:"xmlns,attr"`
}
type params struct {
	XMLNsSoap string `xml:"xmlns,attr"`
	MATNR     string `xml:"MATNR"`
	KUNNR     string `xml:"KUNNR"`
	VTWEG     string `xml:"VTWEG"`
	ZTERM     string `xml:"ZTERM"`
}

func soapCallHandleResponse(ws string, action string, par params, result interface{}) error {
	body, err := StockClientSoap(ws, action, par)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	return nil
}

func StockClientSoap(ws string, action string, par params) ([]byte, error) {
	v := soapRQ{
		XMLNsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		Body: soapBody{
			ZFMWSPRECIOSTOCK{
				XMLNsURN: "urn:sap-com:document:sap:rfc:functions",
				INPUT:    par,
			},
		},
	}
	payload, err := xml.MarshalIndent(v, "", "  ")

	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	req, err := http.NewRequest("POST", ws, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")

	req.Header.Set("SOAPAction", action)
	req.Header.Set("User-Agent", "gowsdl/0.1")

	req.SetBasicAuth("WS_WEB", "ch@ide.2019")

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return bodyBytes, nil
}

/****/

type PRECIOSTOCK struct {
	XMLName xml.Name
	Body    Body
}

type Body struct {
	XMLName     xml.Name
	GetResponse []Response `xml:"ZFM_WS_PRECIOSTOCK"`
}
type Response struct {
	PRECIOSTOCKResponse ItemBody `xml:"ZFM_WS_PRECIOSTOCK"`
}
type ItemBody struct {
	XMLName xml.Name
	Item    itemStok `xml:"item"`
}
type itemStok struct {
	KUNNR string `xml:"KUNNR"`
	MATNR string `xml:"MATNR"`
	VTWEG string `xml:"VTWEG"`

	ZTERM string `xml:"ZTERM"`
	MAKTX string `xml:"MAKTX"`
	KBETR string `xml:"KBETR"`

	DESC1 string `xml:"DESC1"`
	PDSC1 string `xml:"PDSC1"`
	DESC2 string `xml:"DESC2"`

	PDSC2 string `xml:"PDSC2"`
	MNG04 string `xml:"MNG04"`
	PRDTY string `xml:"PRDTY"`
}

func main() {
	ws := "https://srv-hcq-a-qa.industria.chaide.com:1443/sap/bc/srt/rfc/sap/zwsb2b_preciostock/300/zwsb2b_preciostock/zwsb2b_preciostock"
	action := "urn:sap-com:document:sap:rfc:functions:ZWSB2B_PRECIOSTOCK:ZFM_WS_PRECIOSTOCKRequest"
	par := params{
		MATNR: "20004854",
		KUNNR: "",
		VTWEG: "03",
		ZTERM: "DT02",
	}

	Soap, _ := StockClientSoap(ws, action, par)

	fmt.Println(string(Soap))
	res := &PRECIOSTOCK{}
	err := xml.Unmarshal(Soap, res)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(res.Body.GetResponse)*/
}
