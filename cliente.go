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

func ClientSOAPCall(ws, action string, ci string) ([]byte, error) {
	header := `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/">
	<Body>
	  <ZFM_WS_CONSULTACLIENTE xmlns="urn:sap-com:document:sap:rfc:functions">
		<INPUT xmlns="">%s</INPUT>
	  </ZFM_WS_CONSULTACLIENTE>
	</Body>
  </Envelope>`
	headerclient := fmt.Sprintf(header, ci)
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	req, err := http.NewRequest("POST", ws, bytes.NewBuffer([]byte(headerclient)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Set("User-Agent", "gowsdl/0.1")
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

// ClientResponse respuesta del client
type ClientResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	SoapEnv string   `xml:"soap-env,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text                         string `xml:",chardata"`
		ZFMWSCONSULTACLIENTEResponse struct {
			Text   string `xml:",chardata"`
			N0     string `xml:"n0,attr"`
			OUTPUT struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text     string `xml:",chardata"`
					KTOKD    string `xml:"KTOKD"`
					VKORG    string `xml:"VKORG"`
					VTWEG    string `xml:"VTWEG"`
					BZIRK    string `xml:"BZIRK"`
					KDGRP    string `xml:"KDGRP"`
					SPART    string `xml:"SPART"`
					KUNNR    string `xml:"KUNNR"`
					ANRED    string `xml:"ANRED"`
					NAME1    string `xml:"NAME1"`
					MCOD3    string `xml:"MCOD3"`
					LAND1    string `xml:"LAND1"`
					LZONE    string `xml:"LZONE"`
					TELF1    string `xml:"TELF1"`
					TELF2    string `xml:"TELF2"`
					SMTPADDR string `xml:"SMTP_ADDR"`
					TELFX    string `xml:"TELFX"`
					STCD1    string `xml:"STCD1"`
					STCDT    string `xml:"STCDT"`
					FITYP    string `xml:"FITYP"`
					ZTERM    string `xml:"ZTERM"`
					PARVW    string `xml:"PARVW"`
					KUNN2    string `xml:"KUNN2"`
					VTEXT    string `xml:"VTEXT"`
					NAME2    string `xml:"NAME2"`
					ERNAM    string `xml:"ERNAM"`
					ERDAT    string `xml:"ERDAT"`
				} `xml:"item"`
			} `xml:"OUTPUT"`
		} `xml:"ZFM_WS_CONSULTACLIENTEResponse"`
	} `xml:"Body"`
}
