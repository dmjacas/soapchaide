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
func CreateOrder(ws, action string) (*PedidoResponse, error) {

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

	Soap, _ := PedidoSOAPCall(ws, action, data)
	res := &PedidoResponse{}
	err := xml.Unmarshal(Soap, res)
	if err != nil {
		fmt.Println(err)
	}
	return res, err
}
