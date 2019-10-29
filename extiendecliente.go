package soapchaide

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ExtiendeClienteHeader struct {
	XMLName   xml.Name            `xml:"soapenv:Envelope"`
	XMLNsSoap string              `xml:"xmlns:soapenv,attr"`
	XMNsSoap  string              `xml:"xmlns:urn,attr"`
	Body      ExtiendeClienteBody `xml:"soapenv:Body"`
}

type ExtiendeClienteBody struct {
	Zsdb2cExticlie Zsdb2cExticlie `xml:"urn:Zsdb2cExticlie"`
}

type Zsdb2cExticlie struct {
	IvPernr string `xml:"IvPernr"`
	PiKnvv  PiKnvv `xml:"PiKnvv"`
}
type PiKnvv struct {
	Mandt         string `xml:"Mandt,omitempty"`
	Kunnr         string `xml:"Kunnr,omitempty"`
	Vkorg         string `xml:"Vkorg,omitempty"`
	Vtweg         string `xml:"Vtweg,omitempty"`
	Spart         string `xml:"Spart,omitempty"`
	Ernam         string `xml:"Ernam,omitempty"`
	Erdat         string `xml:"Erdat,omitempty"`
	Begru         string `xml:"Begru,omitempty"`
	Loevm         string `xml:"Loevm,omitempty"`
	Versg         string `xml:"Versg,omitempty"`
	Aufsd         string `xml:"Aufsd,omitempty"`
	Kalks         string `xml:"Kalks,omitempty"`
	Kdgrp         string `xml:"Kdgrp,omitempty"`
	Bzirk         string `xml:"Bzirk,omitempty"`
	Konda         string `xml:"Konda,omitempty"`
	Pltyp         string `xml:"Pltyp,omitempty"`
	Awahr         string `xml:"Awahr,omitempty"`
	Inco1         string `xml:"Inco1,omitempty"`
	Inco2         string `xml:"Inco2,omitempty"`
	Lifsd         string `xml:"Lifsd,omitempty"`
	Autlf         string `xml:"Autlf,omitempty"`
	Antlf         string `xml:"Antlf,omitempty"`
	Kztlf         string `xml:"Kztlf,omitempty"`
	Kzazu         string `xml:"Kzazu,omitempty"`
	Chspl         string `xml:"Chspl,omitempty"`
	Lprio         string `xml:"Lprio,omitempty"`
	Eikto         string `xml:"Eikto,omitempty"`
	Vsbed         string `xml:"Vsbed,omitempty"`
	Faksd         string `xml:"Faksd,omitempty"`
	Mrnkz         string `xml:"Mrnkz,omitempty"`
	Perfk         string `xml:"Perfk,omitempty"`
	Perrl         string `xml:"Perrl,omitempty"`
	Kvakz         string `xml:"Kvakz,omitempty"`
	Kvawt         string `xml:"Kvawt,omitempty"`
	Waers         string `xml:"Waers,omitempty"`
	Klabc         string `xml:"Klabc,omitempty"`
	Ktgrd         string `xml:"Ktgrd,omitempty"`
	Zterm         string `xml:"Zterm,omitempty"`
	Vwerk         string `xml:"Vwerk,omitempty"`
	Vkgrp         string `xml:"Vkgrp,omitempty"`
	Vkbur         string `xml:"Vkbur,omitempty"`
	Vsort         string `xml:"Vsort,omitempty"`
	Kvgr1         string `xml:"Kvgr1,omitempty"`
	Kvgr2         string `xml:"Kvgr2,omitempty"`
	Kvgr3         string `xml:"Kvgr3,omitempty"`
	Kvgr4         string `xml:"Kvgr4,omitempty"`
	Kvgr5         string `xml:"Kvgr5,omitempty"`
	Bokre         string `xml:"Bokre,omitempty"`
	Boidt         string `xml:"Boidt,omitempty"`
	Kurst         string `xml:"Kurst,omitempty"`
	Prfre         string `xml:"Prfre,omitempty"`
	Prat1         string `xml:"Prat1,omitempty"`
	Prat2         string `xml:"Prat2,omitempty"`
	Prat3         string `xml:"Prat3,omitempty"`
	Prat4         string `xml:"Prat4,omitempty"`
	Prat5         string `xml:"Prat5,omitempty"`
	Prat6         string `xml:"Prat6,omitempty"`
	Prat7         string `xml:"Prat7,omitempty"`
	Prat8         string `xml:"Prat8,omitempty"`
	Prat9         string `xml:"Prat9,omitempty"`
	Prata         string `xml:"Prata,omitempty"`
	Kabss         string `xml:"Kabss,omitempty"`
	Kkber         string `xml:"Kkber,omitempty"`
	Cassd         string `xml:"Cassd,omitempty"`
	Rdoff         string `xml:"Rdoff,omitempty"`
	Agrel         string `xml:"Agrel,omitempty"`
	Megru         string `xml:"Megru,omitempty"`
	Uebto         string `xml:"Uebto,omitempty"`
	Untto         string `xml:"Untto,omitempty"`
	Uebtk         string `xml:"Uebtk,omitempty"`
	Pvksm         string `xml:"Pvksm,omitempty"`
	Podkz         string `xml:"Podkz,omitempty"`
	Podtg         string `xml:"Podtg,omitempty"`
	Blind         string `xml:"Blind,omitempty"`
	CarrierNotif  string `xml:"CarrierNotif,omitempty"`
	bev1emlgpfand string `xml:"bev1emlgpfand,omitempty"`
	bev1emlgforts string `xml:"bev1emlgforts,omitempty"`
}

func ExtiendeClienteSOA(ws string, action string, data ExtiendeClienteBody) ([]byte, error) {
	v := ExtiendeClienteHeader{
		XMLNsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		XMNsSoap:  "urn:sap-com:document:sap:soap:functions:mc-style",
		Body:      data,
	}
	payload, err := xml.MarshalIndent(v, "", "  ")

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
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", action)
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

type ExtiendeClienteResponse struct {
	XMLName xml.Name `xml:"Envelope"`
}

func main() {
	ws := "http://srv-hcq-a-qa.industria.chaide.com:1080/sap/bc/srt/rfc/sap/zwssdb2cexticlie/300/zwssdb2cexticlie/zwssdb2cexticlie"
	action := "urn:sap-com:document:sap:soap:functions:mc-style:ZWSSDB2CEXTICLIE:Zsdb2cExticlieRequest"
	data := ExtiendeClienteBody{
		Zsdb2cExticlie{
			IvPernr: "",
			PiKnvv: PiKnvv{
				Mandt: "300",
				Kunnr: "0005000020",
				Vkorg: "1000",
				Vtweg: "03",
				Spart: "00",
				Versg: "1",
				Kalks: "1",
				Kdgrp: "05",
				Bzirk: "ZVUIO",
				Awahr: "000",
				Antlf: "0",
				Lprio: "02",
				Vsbed: "01",
				Waers: "USD",
				Ktgrd: "01",
				Zterm: "DT02",
				Vwerk: "1000",
				Vkgrp: "100",
				Vkbur: "100",
			},
		},
	}

	Soap, _ := ExtiendeClienteSOA(ws, action, data)

	res := &ExtiendeClienteResponse{}
	err := xml.Unmarshal(Soap, res)
	fmt.Println(res)
	// fmt.Println(res.Body.GetResponse)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(res.Body.GetResponse)
}
