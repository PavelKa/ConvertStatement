package main

import (
	"reflect"
)

type winstrom struct {
	FakturaV []FakturaVydana `xml:"faktura-vydana"`
}

func newWinstrom() winstrom {
	winstrom := winstrom{}

	return winstrom
}

func (m FakturaVydana) SetByName(fieldName string, value string) Item {
	movementPointer := reflect.ValueOf(&m)

	setFieldValue(movementPointer, fieldName, value)
	return m
}

func (c *winstrom) CreateItem() Item {
	m := FakturaVydana{}
	return m
}
func (c *winstrom) AddItem(item Item) {

	c.FakturaV = append(c.FakturaV, item.(FakturaVydana))

}

type FakturaVydana struct {
	//XMLName xml.Name
	//	id string `xml:"id"`
	Kod      string `xml:"kod"`
	TypDokl  string `xml:"typDokl"`
	VarSym   string `xml:"varSym"`
	NazFirmy string `xml:"nazFirmy"`
	Ulice    string `xml:"ulice"`
	Mesto    string `xml:"mesto"`
	Psc      string `xml:"psc"`
	Ic       string `xml:"ic"`
	Dic      string `xml:"dic"`
	DatVyst  string `xml:"datVyst"`
	DuzpPuv  string `xml:"duzpPuv"`
	DatSplat string `xml:"datSplat"`
	//	datVyst string `xml:"datVyst"`
	SumZklZakl   string `xml:"sumZklZakl"`
	SumDphZakl   string `xml:"sumDphZakl"`
	SumZklCelkem string `xml:"SumZklCelkem"`
	SumZklZaklP  string `xml:"polozkyFaktury>faktura-vydana-polozka>sumZkl"`
}
