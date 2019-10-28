package creacliente

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CreateClientRequest struct {
	XMLName   xml.Name       `xml:"soapenv:Envelope"`
	XMLNsSoap string         `xml:"xmlns:soapenv,attr"`
	XMNsSoap  string         `xml:"xmlns:urn,attr"`
	Body      Zsdb2cCreaclie `xml:"soapenv:Body"`
}

type Zsdb2cCreaclie struct {
	CreateClientBody CreateClientBody `xml:"urn:Zsdb2cCreaclie"`
}

type CreateClientBody struct {
	IBapiaddr1 IBapiaddr1 `xml:"IBapiaddr1,omitempty"`
	IBapiaddr2 IBapiaddr2 `xml:"IBapiaddr2,omitempty"`
	IKna1      IKna1      `xml:"IKna1,omitempty"`
	IKnb1      IKnb1      `xml:"IKnb1,omitempty"`
	IKnvv      IKnvv      `xml:"IKnvv,omitempty"`
	IvPernr    string     `xml:"IvPernr,omitempty"`
	TXknb5     TXknb      `xml:"TXknb5,omitempty"`
	TXknbk     TXknb      `xml:"TXknbk,omitempty"`
	TXknva     TXknb      `xml:"TXknva,omitempty"`
	TXknvi     TXknb      `xml:"TXknvi,omitempty"`
	TXknvk     TXknb      `xml:"TXknvk,omitempty"`
	TXknvp     TXknb      `xml:"TXknvp,omitempty"`
	TYknb5     TXknb      `xml:"TYknb5,omitempty"`
	TYknbk     TXknb      `xml:"TYknbk,omitempty"`
	TYknvi     TXknb      `xml:"TYknvi,omitempty"`
	TYknvk     TXknb      `xml:"TYknvk,omitempty"`
	TYknvp     TXknb      `xml:"TYknvp,omitempty"`
}

type IBapiaddr1 struct {
	AddrNo     string `xml:"AddrNo,omitempty"`
	Formofaddr string `xml:"Formofaddr,omitempty"`
	Name       string `xml:"Name,omitempty"`
	Name2      string `xml:"Name2,omitempty"`
	Name3      string `xml:"Name3,omitempty"`
	Name4      string `xml:"Name4,omitempty"`
	COName     string `xml:"COName,omitempty"`
	City       string `xml:"City,omitempty"`
	District   string `xml:"District,omitempty"`
	CityNo     string `xml:"CityNo,omitempty"`
	PostlCod1  string `xml:"PostlCod1,omitempty"`
	PostlCod2  string `xml:"PostlCod2,omitempty"`
	PostlCod3  string `xml:"PostlCod3,omitempty"`
	PoBox      string `xml:"PoBox,omitempty"`
	PoBoxCit   string `xml:"PoBoxCit,omitempty"`
	DelivDis   string `xml:"DelivDis,omitempty"`
	Street     string `xml:"Street,omitempty"`
	StreetNo   string `xml:"StreetNo,omitempty"`
	StrAbbr    string `xml:"StrAbbr,omitempty"`
	HouseNo    string `xml:"HouseNo,omitempty"`
	StrSuppl1  string `xml:"StrSuppl1,omitempty"`
	StrSuppl2  string `xml:"StrSuppl2,omitempty"`
	Location   string `xml:"Location,omitempty"`
	Building   string `xml:"Building,omitempty"`
	Floor      string `xml:"Floor,omitempty"`
	RoomNo     string `xml:"RoomNo,omitempty"`
	Country    string `xml:"Country,omitempty"`
	Langu      string `xml:"Langu,omitempty"`
	Region     string `xml:"Region,omitempty"`
	Sort1      string `xml:"Sort1,omitempty"`
	Sort2      string `xml:"Sort2,omitempty"`
	TimeZone   string `xml:"TimeZone,omitempty"`
	Taxjurcode string `xml:"Taxjurcode,omitempty"`
	AdrNotes   string `xml:"AdrNotes,omitempty"`
	CommType   string `xml:"CommType,omitempty"`
	Tel1Numbr  string `xml:"Tel1Numbr,omitempty"`
	Tel1Ext    string `xml:"Tel1Ext,omitempty"`
	FaxNumber  string `xml:"FaxNumber,omitempty"`
	FaxExtens  string `xml:"FaxExtens,omitempty"`
	StreetLng  string `xml:"StreetLng,omitempty"`
	DistrctNo  string `xml:"DistrctNo,omitempty"`
	Chckstatus string `xml:"Chckstatus,omitempty"`
	PboxcitNo  string `xml:"PboxcitNo,omitempty"`
	Transpzone string `xml:"Transpzone,omitempty"`
	HouseNo2   string `xml:"HouseNo2,omitempty"`
	EMail      string `xml:"EMail,omitempty"`
	StrSuppl3  string `xml:"StrSuppl3,omitempty"`
	Title      string `xml:"Title,omitempty"`
	Countryiso string `xml:"Countryiso,omitempty"`
	LanguIso   string `xml:"LanguIso,omitempty"`
	BuildLong  string `xml:"BuildLong,omitempty"`
	Regiogroup string `xml:"Regiogroup,omitempty"`
	HomeCity   string `xml:"HomeCity,omitempty"`
	Homecityno string `xml:"Homecityno,omitempty"`
	Pcode1Ext  string `xml:"Pcode1Ext,omitempty"`
	Pcode2Ext  string `xml:"Pcode2Ext,omitempty"`
	Pcode3Ext  string `xml:"Pcode3Ext,omitempty"`
	PoWONo     string `xml:"PoWONo,omitempty"`
	PoBoxReg   string `xml:"PoBoxReg,omitempty"`
	PoboxCtry  string `xml:"PoboxCtry,omitempty"`
	PoCtryiso  string `xml:"PoCtryiso,omitempty"`
	Homepage   string `xml:"Homepage,omitempty"`
	DontUseS   string `xml:"DontUseS,omitempty"`
	DontUseP   string `xml:"DontUseP,omitempty"`
	HouseNo3   string `xml:"HouseNo3,omitempty"`
	LanguCr    string `xml:"LanguCr,omitempty"`
	Langucriso string `xml:"Langucriso,omitempty"`
}
type IBapiaddr2 struct {
	PersNo     string `xml:"PersNo,omitempty"`
	AddrNo     string `xml:"AddrNo,omitempty"`
	TitleP     string `xml:"TitleP,omitempty"`
	Firstname  string `xml:"Firstname,omitempty"`
	Lastname   string `xml:"Lastname,omitempty"`
	BirthName  string `xml:"BirthName,omitempty"`
	Middlename string `xml:"Middlename,omitempty"`
	Secondname string `xml:"Secondname,omitempty"`
	Fullname   string `xml:"Fullname,omitempty"`
	FullnameX  string `xml:"FullnameX,omitempty"`
	TitleAca1  string `xml:"TitleAca1,omitempty"`
	TitleAca2  string `xml:"TitleAca2,omitempty"`
	Prefix1    string `xml:"Prefix1,omitempty"`
	Prefix2    string `xml:"Prefix2,omitempty"`
	TitleSppl  string `xml:"TitleSppl,omitempty"`
	Nickname   string `xml:"Nickname,omitempty"`
	Initials   string `xml:"Initials,omitempty"`
	Nameformat string `xml:"Nameformat,omitempty"`
	Namcountry string `xml:"Namcountry,omitempty"`
	LanguP     string `xml:"LanguP,omitempty"`
	LangupIso  string `xml:"LangupIso,omitempty"`
	Sort1P     string `xml:"Sort1P,omitempty"`
	Sort2P     string `xml:"Sort2P,omitempty"`
	COName     string `xml:"COName,omitempty"`
	City       string `xml:"City,omitempty"`
	District   string `xml:"District,omitempty"`
	CityNo     string `xml:"CityNo,omitempty"`
	DistrctNo  string `xml:"DistrctNo,omitempty"`
	Chckstatus string `xml:"Chckstatus,omitempty"`
	PostlCod1  string `xml:"PostlCod1,omitempty"`
	PostlCod2  string `xml:"PostlCod2,omitempty"`
	PoBox      string `xml:"PoBox,omitempty"`
	PoBoxCit   string `xml:"PoBoxCit,omitempty"`
	PboxcitNo  string `xml:"PboxcitNo,omitempty"`
	DelivDis   string `xml:"DelivDis,omitempty"`
	Transpzone string `xml:"Transpzone,omitempty"`
	Street     string `xml:"Street,omitempty"`
	StreetNo   string `xml:"StreetNo,omitempty"`
	StrAbbr    string `xml:"StrAbbr,omitempty"`
	HouseNo    string `xml:"HouseNo,omitempty"`
	HouseNo2   string `xml:"HouseNo2,omitempty"`
	StrSuppl1  string `xml:"StrSuppl1,omitempty"`
	StrSuppl2  string `xml:"StrSuppl2,omitempty"`
	StrSuppl3  string `xml:"StrSuppl3,omitempty"`
	Location   string `xml:"Location,omitempty"`
	Building   string `xml:"Building,omitempty"`
	Floor      string `xml:"Floor,omitempty"`
	RoomNo     string `xml:"RoomNo,omitempty"`
	Country    string `xml:"Country,omitempty"`
	Countryiso string `xml:"Countryiso,omitempty"`
	Region     string `xml:"Region,omitempty"`
	TimeZone   string `xml:"TimeZone,omitempty"`
	Taxjurcode string `xml:"Taxjurcode,omitempty"`
	AdrNotes   string `xml:"AdrNotes,omitempty"`
	CommType   string `xml:"CommType,omitempty"`
	Tel1Numbr  string `xml:"Tel1Numbr,omitempty"`
	Tel1Ext    string `xml:"Tel1Ext,omitempty"`
	FaxNumber  string `xml:"FaxNumber,omitempty"`
	FaxExtens  string `xml:"FaxExtens,omitempty"`
	EMail      string `xml:"EMail,omitempty"`
	BuildLong  string `xml:"BuildLong,omitempty"`
	Regiogroup string `xml:"Regiogroup,omitempty"`
	HomeCity   string `xml:"HomeCity,omitempty"`
	Homecityno string `xml:"Homecityno,omitempty"`
	Pcode1Ext  string `xml:"Pcode1Ext,omitempty"`
	Pcode2Ext  string `xml:"Pcode2Ext,omitempty"`
	PoWONo     string `xml:"PoWONo,omitempty"`
	PoBoxReg   string `xml:"PoBoxReg,omitempty"`
	PoboxCtry  string `xml:"PoboxCtry,omitempty"`
	PoCtryiso  string `xml:"PoCtryiso,omitempty"`
	Homepage   string `xml:"Homepage,omitempty"`
	DontUseS   string `xml:"DontUseS,omitempty"`
	DontUseP   string `xml:"DontUseP,omitempty"`
	HouseNo3   string `xml:"HouseNo3,omitempty"`
	LanguCrP   string `xml:"LanguCrP,omitempty"`
	Langucpiso string `xml:"Langucpiso,omitempty"`
}
type IKna1 struct {
	Mandt             string `xml:"Mandt,omitempty"`
	Kunnr             string `xml:"Kunnr,omitempty"`
	Land1             string `xml:"Land1,omitempty"`
	Name1             string `xml:"Name1,omitempty"`
	Name2             string `xml:"Name2,omitempty"`
	Ort01             string `xml:"Ort01,omitempty"`
	Pstlz             string `xml:"Pstlz,omitempty"`
	Regio             string `xml:"Regio,omitempty"`
	Sortl             string `xml:"Sortl,omitempty"`
	Stras             string `xml:"Stras,omitempty"`
	Telf1             string `xml:"Telf1,omitempty"`
	Telfx             string `xml:"Telfx,omitempty"`
	Xcpdk             string `xml:"Xcpdk,omitempty"`
	Adrnr             string `xml:"Adrnr,omitempty"`
	Mcod1             string `xml:"Mcod1,omitempty"`
	Mcod2             string `xml:"Mcod2,omitempty"`
	Mcod3             string `xml:"Mcod3,omitempty"`
	Anred             string `xml:"Anred,omitempty"`
	Aufsd             string `xml:"Aufsd,omitempty"`
	Bahne             string `xml:"Bahne,omitempty"`
	Bahns             string `xml:"Bahns,omitempty"`
	Bbbnr             string `xml:"Bbbnr,omitempty"`
	Bbsnr             string `xml:"Bbsnr,omitempty"`
	Begru             string `xml:"Begru,omitempty"`
	Brsch             string `xml:"Brsch,omitempty"`
	Bubkz             string `xml:"Bubkz,omitempty"`
	Datlt             string `xml:"Datlt,omitempty"`
	Erdat             string `xml:"Erdat,omitempty"`
	Ernam             string `xml:"Ernam,omitempty"`
	Exabl             string `xml:"Exabl,omitempty"`
	Faksd             string `xml:"Faksd,omitempty"`
	Fiskn             string `xml:"Fiskn,omitempty"`
	Knazk             string `xml:"Knazk,omitempty"`
	Knrza             string `xml:"Knrza,omitempty"`
	Konzs             string `xml:"Konzs,omitempty"`
	Ktokd             string `xml:"Ktokd,omitempty"`
	Kukla             string `xml:"Kukla,omitempty"`
	Lifnr             string `xml:"Lifnr,omitempty"`
	Lifsd             string `xml:"Lifsd,omitempty"`
	Locco             string `xml:"Locco,omitempty"`
	Loevm             string `xml:"Loevm,omitempty"`
	Name3             string `xml:"Name3,omitempty"`
	Name4             string `xml:"Name4,omitempty"`
	Niels             string `xml:"Niels,omitempty"`
	Ort02             string `xml:"Ort02,omitempty"`
	Pfach             string `xml:"Pfach,omitempty"`
	Pstl2             string `xml:"Pstl2,omitempty"`
	Counc             string `xml:"Counc,omitempty"`
	Cityc             string `xml:"Cityc,omitempty"`
	Rpmkr             string `xml:"Rpmkr,omitempty"`
	Sperr             string `xml:"Sperr,omitempty"`
	Spras             string `xml:"Spras,omitempty"`
	Stcd1             string `xml:"Stcd1,omitempty"`
	Stcd2             string `xml:"Stcd2,omitempty"`
	Stkza             string `xml:"Stkza,omitempty"`
	Stkzu             string `xml:"Stkzu,omitempty"`
	Telbx             string `xml:"Telbx,omitempty"`
	Telf2             string `xml:"Telf2,omitempty"`
	Teltx             string `xml:"Teltx,omitempty"`
	Telx1             string `xml:"Telx1,omitempty"`
	Lzone             string `xml:"Lzone,omitempty"`
	Xzemp             string `xml:"Xzemp,omitempty"`
	Vbund             string `xml:"Vbund,omitempty"`
	Stceg             string `xml:"Stceg,omitempty"`
	Dear1             string `xml:"Dear1,omitempty"`
	Dear2             string `xml:"Dear2,omitempty"`
	Dear3             string `xml:"Dear3,omitempty"`
	Dear4             string `xml:"Dear4,omitempty"`
	Dear5             string `xml:"Dear5,omitempty"`
	Gform             string `xml:"Gform,omitempty"`
	Bran1             string `xml:"Bran1,omitempty"`
	Bran2             string `xml:"Bran2,omitempty"`
	Bran3             string `xml:"Bran3,omitempty"`
	Bran4             string `xml:"Bran4,omitempty"`
	Bran5             string `xml:"Bran5,omitempty"`
	Ekont             string `xml:"Ekont,omitempty"`
	Umsat             string `xml:"Umsat,omitempty"`
	Umjah             string `xml:"Umjah,omitempty"`
	Uwaer             string `xml:"Uwaer,omitempty"`
	Jmzah             string `xml:"Jmzah,omitempty"`
	Jmjah             string `xml:"Jmjah,omitempty"`
	Katr1             string `xml:"Katr1,omitempty"`
	Katr2             string `xml:"Katr2,omitempty"`
	Katr3             string `xml:"Katr3,omitempty"`
	Katr4             string `xml:"Katr4,omitempty"`
	Katr5             string `xml:"Katr5,omitempty"`
	Katr6             string `xml:"Katr6,omitempty"`
	Katr7             string `xml:"Katr7,omitempty"`
	Katr8             string `xml:"Katr8,omitempty"`
	Katr9             string `xml:"Katr9,omitempty"`
	Katr10            string `xml:"Katr10,omitempty"`
	Stkzn             string `xml:"Stkzn,omitempty"`
	Umsa1             string `xml:"Umsa1,omitempty"`
	Txjcd             string `xml:"Txjcd,omitempty"`
	Periv             string `xml:"Periv,omitempty"`
	Abrvw             string `xml:"Abrvw,omitempty"`
	Inspbydebi        string `xml:"Inspbydebi,omitempty"`
	Inspatdebi        string `xml:"Inspatdebi,omitempty"`
	Ktocd             string `xml:"Ktocd,omitempty"`
	Pfort             string `xml:"Pfort,omitempty"`
	Werks             string `xml:"Werks,omitempty"`
	Dtams             string `xml:"Dtams,omitempty"`
	Dtaws             string `xml:"Dtaws,omitempty"`
	Duefl             string `xml:"Duefl,omitempty"`
	Hzuor             string `xml:"Hzuor,omitempty"`
	Sperz             string `xml:"Sperz,omitempty"`
	Etikg             string `xml:"Etikg,omitempty"`
	Civve             string `xml:"Civve,omitempty"`
	Milve             string `xml:"Milve,omitempty"`
	Kdkg1             string `xml:"Kdkg1,omitempty"`
	Kdkg2             string `xml:"Kdkg2,omitempty"`
	Kdkg3             string `xml:"Kdkg3,omitempty"`
	Kdkg4             string `xml:"Kdkg4,omitempty"`
	Kdkg5             string `xml:"Kdkg5,omitempty"`
	Xknza             string `xml:"Xknza,omitempty"`
	Fityp             string `xml:"Fityp,omitempty"`
	Stcdt             string `xml:"Stcdt,omitempty"`
	Stcd3             string `xml:"Stcd3,omitempty"`
	Stcd4             string `xml:"Stcd4,omitempty"`
	Xicms             string `xml:"Xicms,omitempty"`
	Xxipi             string `xml:"Xxipi,omitempty"`
	Xsubt             string `xml:"Xsubt,omitempty"`
	Cfopc             string `xml:"Cfopc,omitempty"`
	Txlw1             string `xml:"Txlw1,omitempty"`
	Txlw2             string `xml:"Txlw2,omitempty"`
	Ccc01             string `xml:"Ccc01,omitempty"`
	Ccc02             string `xml:"Ccc02,omitempty"`
	Ccc03             string `xml:"Ccc03,omitempty"`
	Ccc04             string `xml:"Ccc04,omitempty"`
	Cassd             string `xml:"Cassd,omitempty"`
	Knurl             string `xml:"Knurl,omitempty"`
	J1kfrepre         string `xml:"J1kfrepre,omitempty"`
	J1kftbus          string `xml:"J1kftbus,omitempty"`
	J1kftind          string `xml:"J1kftind,omitempty"`
	Confs             string `xml:"Confs,omitempty"`
	Updat             string `xml:"Updat,omitempty"`
	Uptim             string `xml:"Uptim,omitempty"`
	Nodel             string `xml:"Nodel,omitempty"`
	Dear6             string `xml:"Dear6,omitempty"`
	Psoo5Psoo5rPalhgt string `xml:"_-vso_-rPalhgt,omitempty"`
	Psoo5rPalUl       string `xml:"_-vso_-rPalUl,omitempty"`
	Psoo5rPkMat       string `xml:"_-vso_-rPkMat,omitempty"`
	Psoo5rMatpal      string `xml:"_-vso_-rMatpal,omitempty"`
	Psoo5rINoLyr      string `xml:"_-vso_-rINoLyr,omitempty"`
	Psoo5rOneMat      string `xml:"_-vso_-rOneMat,omitempty"`
	Psoo5rOneSort     string `xml:"_-vso_-rOneSort,omitempty"`
	Psoo5rUldSide     string `xml:"_-vso_-rUldSide,omitempty"`
	Psoo5rLoadPref    string `xml:"_-vso_-rLoadPref,omitempty"`
	Psoo5rDpoint      string `xml:"_-vso_-rDpoint,omitempty"`
	Alc               string `xml:"Alc,omitempty"`
	PmtOffice         string `xml:"PmtOffice,omitempty"`
	Psofg             string `xml:"Psofg,omitempty"`
	Psois             string `xml:"Psois,omitempty"`
	Pson1             string `xml:"Pson1,omitempty"`
	Pson2             string `xml:"Pson2,omitempty"`
	Pson3             string `xml:"Pson3,omitempty"`
	Psovn             string `xml:"Psovn,omitempty"`
	Psotl             string `xml:"Psotl,omitempty"`
	Psohs             string `xml:"Psohs,omitempty"`
	Psost             string `xml:"Psost,omitempty"`
	Psoo1             string `xml:"Psoo1,omitempty"`
	Psoo2             string `xml:"Psoo2,omitempty"`
	Psoo3             string `xml:"Psoo3,omitempty"`
	Psoo4             string `xml:"Psoo4,omitempty"`
	Psoo5             string `xml:"Psoo5,omitempty"`
}

type IKnb1 struct {
	Mandt     string `xml:"Mandt,omitempty"`
	Kunnr     string `xml:"Kunnr,omitempty"`
	Bukrs     string `xml:"Bukrs,omitempty"`
	Pernr     string `xml:"Pernr,omitempty"`
	Erdat     string `xml:"Erdat,omitempty"`
	Ernam     string `xml:"Ernam,omitempty"`
	Sperr     string `xml:"Sperr,omitempty"`
	Loevm     string `xml:"Loevm,omitempty"`
	Zuawa     string `xml:"Zuawa,omitempty"`
	Busab     string `xml:"Busab,omitempty"`
	Akont     string `xml:"Akont,omitempty"`
	Begru     string `xml:"Begru,omitempty"`
	Knrze     string `xml:"Knrze,omitempty"`
	Knrzb     string `xml:"Knrzb,omitempty"`
	Zamim     string `xml:"Zamim,omitempty"`
	Zamiv     string `xml:"Zamiv,omitempty"`
	Zamir     string `xml:"Zamir,omitempty"`
	Zamib     string `xml:"Zamib,omitempty"`
	Zamio     string `xml:"Zamio,omitempty"`
	Zwels     string `xml:"Zwels,omitempty"`
	Xverr     string `xml:"Xverr,omitempty"`
	Zahls     string `xml:"Zahls,omitempty"`
	Zterm     string `xml:"Zterm,omitempty"`
	Wakon     string `xml:"Wakon,omitempty"`
	Vzskz     string `xml:"Vzskz,omitempty"`
	Zindt     string `xml:"Zindt,omitempty"`
	Zinrt     string `xml:"Zinrt,omitempty"`
	Eikto     string `xml:"Eikto,omitempty"`
	Zsabe     string `xml:"Zsabe,omitempty"`
	Kverm     string `xml:"Kverm,omitempty"`
	Fdgrv     string `xml:"Fdgrv,omitempty"`
	Vrbkz     string `xml:"Vrbkz,omitempty"`
	Vlibb     string `xml:"Vlibb,omitempty"`
	Vrszl     string `xml:"Vrszl,omitempty"`
	Vrspr     string `xml:"Vrspr,omitempty"`
	Vrsnr     string `xml:"Vrsnr,omitempty"`
	Verdt     string `xml:"Verdt,omitempty"`
	Perkz     string `xml:"Perkz,omitempty"`
	Xdezv     string `xml:"Xdezv,omitempty"`
	Xausz     string `xml:"Xausz,omitempty"`
	Webtr     string `xml:"Webtr,omitempty"`
	Remit     string `xml:"Remit,omitempty"`
	Datlz     string `xml:"Datlz,omitempty"`
	Xzver     string `xml:"Xzver,omitempty"`
	Togru     string `xml:"Togru,omitempty"`
	Kultg     string `xml:"Kultg,omitempty"`
	Hbkid     string `xml:"Hbkid,omitempty"`
	Xpore     string `xml:"Xpore,omitempty"`
	Blnkz     string `xml:"Blnkz,omitempty"`
	Altkn     string `xml:"Altkn,omitempty"`
	Zgrup     string `xml:"Zgrup,omitempty"`
	Urlid     string `xml:"Urlid,omitempty"`
	Mgrup     string `xml:"Mgrup,omitempty"`
	Lockb     string `xml:"Lockbv,omitempty"`
	Uzawe     string `xml:"Uzawe,omitempty"`
	Ekvbd     string `xml:"Ekvbd,omitempty"`
	Sregl     string `xml:"Sregl,omitempty"`
	Xedip     string `xml:"Xedip,omitempty"`
	Frgrp     string `xml:"Frgrp,omitempty"`
	Vrsdg     string `xml:"Vrsdgs,omitempty"`
	Tlfxs     string `xml:"Tlfxs,omitempty"`
	Intad     string `xml:"Intad,omitempty"`
	Xknzb     string `xml:"Xknzb,omitempty"`
	Guzte     string `xml:"Guzte,omitempty"`
	Gricd     string `xml:"Gricd,omitempty"`
	Gridt     string `xml:"Gridt,omitempty"`
	Wbrsl     string `xml:"Wbrsl,omitempty"`
	Confs     string `xml:"Confs,omitempty"`
	Updat     string `xml:"Updat,omitempty"`
	Uptim     string `xml:"Uptim,omitempty"`
	Nodel     string `xml:"Nodel,omitempty"`
	Tlfns     string `xml:"Tlfns,omitempty"`
	CessionKz string `xml:"CessionKz,omitempty"`
	Gmvkzd    string `xml:"Gmvkzd,omitempty"`
}

type IKnvv struct {
	Mandt         string `xml:"Mandt,omitempty"`
	Kunnr         string `xml:"Kunnr,omitempty"`
	Vkorg         string `xml:"Vkorg,omitempty"`
	Vtweg         string `xml:"Vtweg,omitempty"`
	Spart         string `xml:"Spart,omitempty"`
	IKnvv         string `xml:"IKnvv,omitempty"`
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
	Bev1Emlgpfand string `xml:"_-bev1_-emlgpfand,omitempty"`
	Bev1Emlgforts string `xml:"_-bev1_-emlgforts,omitempty"`
}
type TXknb struct {
	Item CreateClientItem `xml:"item"`
}

type CreateClientItem struct {
	Aland   string `xml:"Aland,omitempty"`
	Tatyp   string `xml:"Tatyp,omitempty"`
	Taxkd   string `xml:"Taxkd,omitempty"`
	Mandt   string `xml:"Mandt,omitempty"`
	Parnr   string `xml:"Parnr,omitempty"`
	Kunnr   string `xml:"Kunnr,omitempty"`
	Namev   string `xml:"Namev,omitempty"`
	Name1   string `xml:"Name1,omitempty"`
	Ort01   string `xml:"Ort01,omitempty"`
	Adrnd   string `xml:"Adrnd,omitempty"`
	Adrnp   string `xml:"Adrnp,omitempty"`
	Abtpa   string `xml:"Abtpa,omitempty"`
	Abtnr   string `xml:"Abtnr,omitempty"`
	Uepar   string `xml:"Uepar,omitempty"`
	Telf1   string `xml:"Telf1,omitempty"`
	Anred   string `xml:"Anred,omitempty"`
	Pafkt   string `xml:"Pafkt,omitempty"`
	Parvo   string `xml:"Parvo,omitempty"`
	Pavip   string `xml:"Pavip,omitempty"`
	Parge   string `xml:"Parge,omitempty"`
	Parla   string `xml:"Parla,omitempty"`
	Gbdat   string `xml:"Gbdat,omitempty"`
	Vrtnr   string `xml:"Vrtnr,omitempty"`
	Bryth   string `xml:"Bryth,omitempty"`
	Akver   string `xml:"Akver,omitempty"`
	Nmail   string `xml:"Nmail,omitempty"`
	Parau   string `xml:"Parau,omitempty"`
	Parh1   string `xml:"Parh1,omitempty"`
	Parh2   string `xml:"Parh2,omitempty"`
	Parh3   string `xml:"Parh3,omitempty"`
	Parh4   string `xml:"Parh4,omitempty"`
	Parh5   string `xml:"Parh5,omitempty"`
	Moab1   string `xml:"Moab1,omitempty"`
	Mobi1   string `xml:"Mobi1,omitempty"`
	Moab2   string `xml:"Moab2,omitempty"`
	Mobi2   string `xml:"Mobi2,omitempty"`
	Diab1   string `xml:"Diab1,omitempty"`
	Dibi1   string `xml:"Dibi1,omitempty"`
	Diab2   string `xml:"Diab2,omitempty"`
	Dibi2   string `xml:"Dibi2,omitempty"`
	Miab1   string `xml:"Miab1,omitempty"`
	Mibi1   string `xml:"Mibi1,omitempty"`
	Miab2   string `xml:"Miab2,omitempty"`
	Mibi2   string `xml:"Mibi2,omitempty"`
	Doab1   string `xml:"Doab1,omitempty"`
	Dobi1   string `xml:"Dobi1,omitempty"`
	Doab2   string `xml:"Doab2,omitempty"`
	Dobi2   string `xml:"Dobi2,omitempty"`
	Frab1   string `xml:"Frab1,omitempty"`
	Frbi1   string `xml:"Frbi1,omitempty"`
	Frab2   string `xml:"Frab2,omitempty"`
	Frbi2   string `xml:"Frbi2,omitempty"`
	Sabi1   string `xml:"Sabi1,omitempty"`
	Saab2   string `xml:"Saab2,omitempty"`
	Sabi2   string `xml:"Sabi2,omitempty"`
	Soab1   string `xml:"Soab1,omitempty"`
	Sobi1   string `xml:"Sobi1,omitempty"`
	Soab2   string `xml:"Soab2,omitempty"`
	Sobi2   string `xml:"Sobi2,omitempty"`
	Pakn1   string `xml:"Pakn1,omitempty"`
	Pakn2   string `xml:"Pakn2,omitempty"`
	Pakn3   string `xml:"Pakn3,omitempty"`
	Pakn4   string `xml:"Pakn4,omitempty"`
	Pakn5   string `xml:"Pakn5,omitempty"`
	Sortl   string `xml:"Sortl,omitempty"`
	Famst   string `xml:"Famst,omitempty"`
	Spnam   string `xml:"Spnam,omitempty"`
	TitelAp string `xml:"TitelAp,omitempty"`
	Erdat   string `xml:"Erdat,omitempty"`
	Ernam   string `xml:"Ernam,omitempty"`
	Duefl   string `xml:"Duefl,omitempty"`
	Lifnr   string `xml:"Lifnr,omitempty"`
	Loevm   string `xml:"Loevm,omitempty"`
	Kzherk  string `xml:"Kzherk,omitempty"`
	Adrnp2  string `xml:"Adrnp2,omitempty"`
	Prsnr   string `xml:"Prsnr,omitempty"`
	Kz      string `xml:"Kz,omitempty"`
}

func soapCallHandleResponse(ws, action string, par Zsdb2cCreaclie, result interface{}) error {
	body, err := CreateClientSOAP(ws, action, par)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	return nil
}

func CreateClientSOAP(ws string, action string, data Zsdb2cCreaclie) ([]byte, error) {
	v := CreateClientRequest{
		XMLNsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		XMNsSoap:  "urn:sap-com:document:sap:soap:functions:mc-style",
		Body:      data,
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
	//fmt.Println(string(payload))
	///	return nil, nil

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
	fmt.Println(string(bodyBytes))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return bodyBytes, nil
}

/****/

type ClientResponse struct {
	XMLName xml.Name     `xml:"Envelope"`
	Body    BodyResponse `xml:"Body"`
}
type HeaderResponse struct {
}
type BodyResponse struct {
	XMLName     xml.Name
	GetResponse []Response `xml:"ZFM_WS_CONSULTACLIENTEResponse"`
}
type Response struct {
	OUTPUT ItemBody `xml:"OUTPUT"`
}
type ItemBody struct {
	Item itemStok `xml:"item"`
}
type itemStok struct {
	KTOKD string `xml:"KTOKD"`
	VKORG string `xml:"VKORG"`
	VTWEG string `xml:"VTWEG"`

	BZIRK string `xml:"BZIRK"`
	KDGRP string `xml:"KDGRP"`
	SPART string `xml:"SPART"`
}

func main() {
	ws := "http://srv-hcq-a-qa.industria.chaide.com:1080/sap/bc/srt/rfc/sap/zwssdb2ccreacli/300/zwssdb2ccreacli/zwssdb2ccreacli"
	action := "urn:sap-com:document:sap:soap:functions:mc-style:ZWSSDB2CCREACLI:Zsdb2cCreaclieRequest"

	ci := Zsdb2cCreaclie{
		CreateClientBody{
			IBapiaddr2: IBapiaddr2{
				TitleP:    "Señora",
				Firstname: "HIDALGO ABRIL",
				Lastname:  "Lastname",
				LanguP:    "S",
				Sort1P:    "1716032774",
				City:      "QUITO",
				Street:    "VIA INTERVALLES KM 2.5 URB PALL S/N",
				StrSuppl1: "CASA 137",
				StrSuppl2: "PUERTA NEGRA",
				Country:   "EC",
				Region:    "17",
				Tel1Numbr: "023989103",
				EMail:     "derix.gomez@chaideychaide.com",
			},
			IKna1: IKna1{
				Mandt: "300",
				Land1: "EC",
				Name1: "ALEX PAUL HIDALGO ABRIL",
				Name2: "ALEX PAUL HIDALGO ABRIL",
				Regio: "17",
				Sortl: "1716032774",
				Stras: "VIA INTERVALLES KM 2.5 URB PALL S/N",
				Telf1: "023989103",
				Mcod3: "QUITO",
				Anred: "Señora",
				Bbbnr: "0000000",
				Bbsnr: "00000",
				Bubkz: "0",
				Ktokd: "BP01",
				Kukla: "04",
				Ort02: "SMS+CORREO",
				Spras: "S",
				Stcd1: "1716032774",
				Stkzu: "X",
				Telf2: "0990000002",
				Lzone: "UIO_PLANTA",
				Stkzn: "X",
				Duefl: "X",
				Kdkg1: "04",
				Fityp: "PN",
				Stcdt: "05",
			},
			IKnb1: IKnb1{
				Mandt: "300",
				Bukrs: "1000",
				Pernr: "00000000",
				Busab: "15",
				Akont: "1301010101",
				Zwels: "CDEFT",
				Zterm: "DT02",
				Fdgrv: "E9",
				Vlibb: "0.0",
				Vrszl: "0",
				Vrspr: "0",
				Xausz: "1",
				Webtr: "0.0",
				Xzver: "X",
				Kultg: "0",
			},
			IKnvv: IKnvv{
				Mandt: "300",
				Vkorg: "1000",
				Vtweg: "03",
				Spart: "00",
				Versg: "1",
				Kalks: "1",
				Kdgrp: "05",
				Bzirk: "ZVUIO",
				Awahr: "100",
				Antlf: "0",
				Kzazu: "X",
				Lprio: "02",
				Vsbed: "01",
				Kvawt: "0",
				Waers: "USD",
				Ktgrd: "01",
				Zterm: "DT02",
				Vwerk: "1000",
				Vkgrp: "100",
				Vkbur: "100",
				Uebto: "0",
				Untto: "0",
			},
			TXknva: TXknb{
				Item: CreateClientItem{},
			},
			TXknvi: TXknb{
				Item: CreateClientItem{
					Mandt: "300",
					Aland: "EC",
					Tatyp: "MWST",
					Taxkd: "1",
				},
			},
			TXknvk: TXknb{
				Item: CreateClientItem{},
			},
			TXknvp: TXknb{
				Item: CreateClientItem{},
			},
			TYknb5: TXknb{
				Item: CreateClientItem{},
			},
			TYknbk: TXknb{
				Item: CreateClientItem{},
			},
			TYknvi: TXknb{
				Item: CreateClientItem{},
			},
			TYknvk: TXknb{
				Item: CreateClientItem{},
			},
			TYknvp: TXknb{
				Item: CreateClientItem{},
			},
		},
	}

	Soap, _ := CreateClientSOAP(ws, action, ci)

	res := &ClientResponse{}
	err := xml.Unmarshal(Soap, res)
	fmt.Println(res)
	// fmt.Println(res.Body.GetResponse)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(res.Body.GetResponse)
}
