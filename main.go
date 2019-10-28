package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type PedidoRequest struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	URN     string   `xml:"urn,attr"`
	Body    struct {
		Text         string `xml:",chardata"`
		Zsdb2Creaped struct {
			Text      string `xml:",chardata"`
			Auart     string `xml:"Auart"`
			Augru     string `xml:"Augru"`
			Erdat     string `xml:"Erdat"`
			Kunnr     string `xml:"Kunnr"`
			Kunnr1    string `xml:"Kunnr1"`
			PurchNoC  string `xml:"PurchNoC"`
			PurchNoS  string `xml:"PurchNoS"`
			Spart     string `xml:"Spart"`
			Vkogr     string `xml:"Vkogr"`
			Vtweg     string `xml:"Vtweg"`
			Zterm     string `xml:"Zterm"`
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
		} `xml:"Zsdb2Creaped"`
	} `xml:"Body"`
}

//PedidoSOAPCall llamado a servicio de crear pedido
func PedidoSOAPCall(ws string, action string, par PedidoRequest) ([]byte, error) {

	payload, err := xml.MarshalIndent(par, "", "  ")
	fmt.Println(string(payload))
	return nil, nil
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
	fmt.Println(string(bodyBytes))

	defer response.Body.Close()
	return bodyBytes, nil
}

/**Response*/

type PedidoResponse struct {
	XMLName xml.Name     `xml:"Envelope"`
	Text    string       `xml:",chardata"`
	Soapenv string       `xml:"soapenv,attr"`
	URN     string       `xml:"urn,attr"`
	Body    PurchaseBody `xml:"Body"`
}

type PurchaseBody struct {
	Text         string       `xml:",chardata"`
	Zsdb2Creaped Zsdb2Creaped `xml:"Zsdb2Creaped"`
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
	Zterm     string    `xml:"Zterm"`
	Zsdb2Item Zsdb2Item `xml:"Zsdb2Item"`
}
type Zsdb2Item struct {
	Text string         `xml:",chardata"`
	Item []PurchaseItem `xml:"item"`
}
type PurchaseItem struct {
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

func main() {
	/*
		ws := "http://srv-hcq-a-qa.industria.chaide.com:1080/sap/bc/srt/rfc/sap/zwssdb2bcreaped/300/zwssdb2bcreaped/zwssdb2bcreaped"
		action := "urn:sap-com:document:sap:soap:functions:mc-style:ZWSSDB2BCREAPED:Zsdb2CreapedRequest"
		items1 := subItem{
			Mandt:  "300",
			Posnr:  "1",
			Matnr:  "20001679",
			Kwmeng: "1",
			Pstyv:  "TAN",
			Kunnr:  "0005000020",
			Etdat:  "10102019",
			Kwert:  "",
			Abgru:  "",
		}
		items2 := subItem{
			Mandt:  "300",
			Posnr:  "1",
			Matnr:  "20001679",
			Kwmeng: "1",
			Pstyv:  "TAN",
			Kunnr:  "0005000020",
			Etdat:  "10102019",
			Kwert:  "",
			Abgru:  "",
		}
		subItems := []subItem{items1, items2}

		Zsdb2Item := itemsPedido{Zsdb2Item: subItems}
		data := Pedido{
			Auart:     "ZTA1",
			Augru:     "001",
			Erdat:     "10102019",
			Kunnr:     "0005000020",
			Kunnr1:    "0005000020",
			PurchNoC:  "",
			PurchNoS:  "",
			Spart:     "00",
			Vkogr:     "1000",
			Vtweg:     "03",
			Zterm:     "DT02",
			Zsdb2Item: Zsdb2Item,
		}


		v := PedidoRequest{
			XMLNsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
			XMLNsUrn:  "urn:sap-com:document:sap:soap:functions:mc-style",
			Body: soapPedidoBody{
				Zsdb2Creaped: par,
			},
		}

	*/
	data := PedidoRequest{
		Soapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		URN:     "urn:sap-com:document:sap:soap:functions:mc-style",
		Body: PurchaseBody{
			Zsdb2Creaped: Zsdb2Creaped{
				Auart:    "ZTA1",
				Augru:    "001",
				Erdat:    "10102019",
				Kunnr:    "0005000020",
				Kunnr1:   "0005000020",
				PurchNoC: "",
				PurchNoS: "",
				Spart:    "00",
				Vkogr:    "1000",
				Vtweg:    "03",
				Zterm:    "DT02",
			},
		},
	}
	fmt.Println(data)
	/*
		Soap, _ := PedidoSOAPCall(ws, action, data)
		res := &PedidoResponse{}
		err := xml.Unmarshal(Soap, res)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res.Body.Zsdb2CreapedResponse)*/
}
