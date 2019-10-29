package soapchaide

import (
	"encoding/xml"
	"fmt"
)

//GetStock informacion product stok
func GetStock(ws, action, code string) (*StockResponse, error) {
	par := params{
		MATNR: code,
		KUNNR: "",
		VTWEG: "03",
		ZTERM: "DT02",
	}

	Soap, _ := StockClientSoap(ws, action, par)
	res := &StockResponse{}

	err := xml.Unmarshal(Soap, res)
	if err != nil {
		fmt.Println(err)
	}

	return res, nil
}

//GetClientInfo informacion del cliente
func GetClientInfo(ws, action, ci string) (*ClientResponse, error) {

	Soap, _ := ClientSOAPCall(ws, action, ci)
	res := &ClientResponse{}
	err := xml.Unmarshal(Soap, res)
	if err != nil {
		fmt.Println(err)
	}
	return res, err
}

//CreateOrder create order
func CreateOrder(ws string, action string, data Zsdb2Creaped) (*PedidoResponse, error) {
	Soap, _ := PedidoSOAPCall(ws, action, data)
	res := &PedidoResponse{}
	err := xml.Unmarshal(Soap, res)
	if err != nil {
		fmt.Println(err)
	}
	return res, err
}

//CreateClient create client
func CreateClient(ws string, action string, ci Zsdb2cCreaclie) (*CreateClientResponse, error) {

	Soap, _ := CreateClientSOAP(ws, action, ci)

	res := &CreateClientResponse{}
	err := xml.Unmarshal(Soap, res)
	if err != nil {
		fmt.Println(err)
	}
	return res, nil
}
