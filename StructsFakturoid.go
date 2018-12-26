package main

type winstrom struct {
	FakturaV []FakturaVydana `xml:"faktura-vydana"`
}

func newWinstrom() winstrom {
	winstrom := winstrom{}

	return winstrom
}

func (c *winstrom) CreateItem() interface{} {
	m := FakturaVydana{}
	//c.FakturaV = append(c.FakturaV, m)

	return &m
}
func (c *winstrom) AddItem(item interface{}) {
	fv := (item).(*FakturaVydana)
	c.FakturaV = append(c.FakturaV, *fv)

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
