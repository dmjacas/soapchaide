package soapchaide

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"
)

type PedidoRequest struct {
	XMLName xml.Name          `xml:"soapenv:Envelope"`
	Text    string            `xml:",chardata"`
	Soapenv string            `xml:"xmlns:soapenv,attr"`
	URN     string            `xml:"xmlns:urn,attr"`
	Header  string            `xml:"soapenv:Header"`
	Body    BodyPedidoRequest `xml:"soapenv:Body"`
}

type BodyPedidoRequest struct {
	Zsdb2Creaped Zsdb2Creaped `xml:"urn:Zsdb2Creaped"`
}

type Zsdb2Creaped struct {
	Text      string    `xml:",chardata"`
	Auart     string    `xml:"Auart"`
	Augru     string    `xml:"Augru"`
	Erdat     string    `xml:"Erdat"`
	Kunnr     string    `xml:"Kunnr"`
	Kunnr1    string    `xml:"Kunnr1"`
	PurchNoC  string    `xml:"PurchNoC"`
	PurchNoS  string    `xml:"PurchNoS"`
	Spart     string    `xml:"Spart"`
	Vkogr     string    `xml:"Vkogr"`
	Vtweg     string    `xml:"Vtweg"`
	Zsdb2Item Zsdb2Item `xml:"Zsdb2Item"`
	Zterm     string    `xml:"Zterm"`
}

type Zsdb2Item struct {
	Text string              `xml:",chardata"`
	Item []PedidoItemRequest `xml:"item"`
}

type PedidoItemRequest struct {
	Text   string `xml:",chardata"`
	Mandt  string `xml:"Mandt"`
	Posnr  string `xml:"Posnr"`
	Matnr  string `xml:"Matnr"`
	Kwmeng string `xml:"Kwmeng"`
	Pstyv  string `xml:"Pstyv"`
	Kunnr  string `xml:"Kunnr"`
	Etdat  string `xml:"Etdat"`
	Kwert  string `xml:"Kwert"`
	Abgru  string `xml:"Abgru"`
}

//PedidoSOAPCall llamado a servicio de crear pedido
func PedidoSOAPCall(ws string, action string, par Zsdb2Creaped) ([]byte, error) {
	v := PedidoRequest{
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		URN:     "urn:sap-com:document:sap:soap:functions:mc-style",
		Body: BodyPedidoRequest{
			Zsdb2Creaped: par,
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

type PedidoResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SoapEnv string   `xml:"soap-env,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text                 string `xml:",chardata"`
		Zsdb2CreapedResponse struct {
			Text      string `xml:",chardata"`
			N0        string `xml:"n0,attr"`
			Kwert     string `xml:"Kwert"`
			Kwert1    string `xml:"Kwert1"`
			Kwert2    string `xml:"Kwert2"`
			Kwert3    string `xml:"Kwert3"`
			TextM     string `xml:"TextM"`
			TypeM     string `xml:"TypeM"`
			Vbeln     string `xml:"Vbeln"`
			Zsdb2Item struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text   string `xml:",chardata"`
					Mandt  string `xml:"Mandt"`
					Posnr  string `xml:"Posnr"`
					Matnr  string `xml:"Matnr"`
					Kwmeng string `xml:"Kwmeng"`
					Pstyv  string `xml:"Pstyv"`
					Kunnr  string `xml:"Kunnr"`
					Etdat  string `xml:"Etdat"`
					Kwert  string `xml:"Kwert"`
					Abgru  string `xml:"Abgru"`
				} `xml:"item"`
			} `xml:"Zsdb2Item"`
		} `xml:"Zsdb2CreapedResponse"`
	} `xml:"Body"`
}
