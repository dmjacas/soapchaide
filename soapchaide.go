package soapchaide

import (
	"encoding/xml"
	"fmt"
)

func getClientInfo(ws, action, ci) (ClientResponse, error) {
	/*
		ws := "https://srv-hcq-a-qa.industria.chaide.com:1443/sap/bc/srt/rfc/sap/zwsb2b_consultacliente/300/zfm_ws_consultacliente/zfm_ws_consultacliente"
		action := "urn:sap-com:document:sap:rfc:functions:ZWSB2B_CONSULTACLIENTE:ZFM_WS_CONSULTACLIENTERequest"
		ci := "1716032774"
	*/
	Soap, _ := ClientSOAPCall(ws, action, ci)
	res := &ClientResponse{}
	err := xml.Unmarshal(Soap, res)
	fmt.Println(res.Body.GetResponse)
	if err != nil {
		fmt.Println(err)
	}
	return res, err
}
