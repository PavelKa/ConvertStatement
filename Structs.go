package main

type AccountMovements struct {
	official               string `xml:"Amount,attr"`
	statemDebitTotal       string `xml:"StatemDebitTotal,attr"`
	statemCreditTotal      string `xml:"StatemCreditTotal,attr"`
	statemTransactionCount string `xml:"StatemTransactionCount,attr"`
	statemDebitCount       string `xml:"StatemDebitCount"`
	statemCreditCount      string `xml:"StatemCreditCount,attr"`
	Movement               []Movement
}

func NewAccountMovements() AccountMovements {
	accountMovements := AccountMovements{}
	accountMovements.official = "N"
	accountMovements.statemDebitTotal = "0"
	return accountMovements
}

func (c *AccountMovements) AddMovement(movement Movement) {
	c.Movement = append(c.Movement, movement)
}

type Movement struct {
	//XMLName xml.Name
	Amount           string `xml:"Amount,attr"`
	Direction        string `xml:"Direction,attr"`
	PostingDate      string `xml:"PostingDate,attr"`
	ItemNo           string `xml:"ItemNo,attr"`
	PartnerAccNo     string
	PartnerAccBank   string
	PartnerAccName   string
	ValueDate        string
	ChargesAmount    string
	ChargesCcy       string
	MovementTypeText string
	BankCode         string
	BankCountryID    string
	BankName         string
	AccNoID          string
	AccNoCC          string
	AccName          string
	AccCcy           string
	AccCcyText       string
	Statistics1      string
	Statistics2      string
	Statistics3      string
	Statistics4      string
	Description1     string
}

func NewMovement() Movement {
	movement := Movement{}

	return movement
}
