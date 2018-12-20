package main

import (
	"strconv"
	"strings"
	"time"
)

type MovementFieldFunction interface{}

func getRecItem(rec []string, i int) string {
	if i >= 0 {
		return rec[i]
	} else {
		return ""
	}
}

func getAccBank(rec []string, i int) string {
	nf := strings.Split(rec[i], "/")
	if len(nf) > 1 {
		return nf[1]
	} else {
		return ""
	}

}

func getDate(rec []string, i int) string {
	t1, e := time.Parse(
		"2.01.2006",
		strings.Split(rec[i], " ")[0])
	check(e)

	return t1.Format("20060102")

}

func getAmount(rec []string, i int) string {

	nf := strings.Replace(rec[i], " ", "", -1)

	return nf

}

func getAccP2(rec []string, i int) string {
	nf := strings.Split(rec[i], "/")

	return nf[0]

}
func getDirection(rec []string, i int) string {
	if i >= 0 {
		nf := strings.Replace(rec[i], " ", "", -1)
		nf = strings.Replace(nf, ",", ".", -1)
		f, err := strconv.ParseFloat(nf, 64)
		check(err)
		if f < 0 {
			return "C"
		} else {
			return "D"
		}
	} else {
		return "U"
	}
}

type MovementImpl struct {
	index    int
	function MovementFieldFunction
}

type MovementFields struct {
	Amount           MovementImpl
	Direction        MovementImpl
	PostingDate      MovementImpl
	ItemNo           MovementImpl
	PartnerAccNo     MovementImpl
	PartnerAccBank   MovementImpl
	PartnerAccName   MovementImpl
	ValueDate        MovementImpl
	ChargesAmount    MovementImpl
	ChargesCcy       MovementImpl
	MovementTypeText MovementImpl
	BankCode         MovementImpl
	BankCountryID    MovementImpl
	BankName         MovementImpl
	AccNoID          MovementImpl
	AccNoCC          MovementImpl
	AccName          MovementImpl
	AccCcy           MovementImpl
	AccCcyText       MovementImpl
	Statistics1      MovementImpl
	Statistics2      MovementImpl
	Statistics3      MovementImpl
	Statistics4      MovementImpl
	Description1     MovementImpl
}

//0	Datum provedení
//1	Datum zaúčtování
//2	Číslo účtu
//3	Název účtu
//4	Kategorie transakce
//5	Číslo protiúčtu
//6	Název protiúčtu
//7	Typ transakce
//8	Zpráva
//9	Poznámka
//10	VS
//11	KS
//12	SS
//13	Zaúčtovaná částka
//14	Měna účtu
//15	Původní částka a měna
//16	Původní částka a měna
//17	Poplatek
//18	Id transakce

func newMovementImpl(field int) MovementImpl {
	mvf := MovementImpl{}
	mvf.function = getRecItem
	mvf.index = field
	return mvf

}
func newMovementImplF(field int, f MovementFieldFunction) MovementImpl {
	mvf := MovementImpl{}
	mvf.function = f
	mvf.index = field
	return mvf

}
func newMovementFieldRsB() MovementFields {
	movementFields := MovementFields{}
	movementFields.Amount = newMovementImplF(13, getAmount)

	movementFields.Direction = newMovementImplF(13, getDirection)

	movementFields.PostingDate = newMovementImplF(0, getDate)
	movementFields.ItemNo = newMovementImpl(18)
	movementFields.PartnerAccNo = newMovementImplF(5, getAccP2)
	movementFields.PartnerAccBank = newMovementImplF(5, getAccBank)
	movementFields.PartnerAccName = newMovementImpl(6)
	movementFields.ValueDate = newMovementImplF(1, getDate)
	movementFields.ChargesAmount = newMovementImplF(17, getAmount)
	movementFields.ChargesCcy = newMovementImpl(17)
	movementFields.MovementTypeText = newMovementImpl(7)
	movementFields.BankCode = newMovementImplF(2, getAccBank)
	movementFields.BankCountryID = newMovementImpl(-1)
	movementFields.BankName = newMovementImpl(-1)
	movementFields.AccNoID = newMovementImplF(2, getAccP2)
	movementFields.AccNoCC = newMovementImpl(-1)
	movementFields.AccName = newMovementImpl(3)
	movementFields.AccCcy = newMovementImpl(14)
	movementFields.AccCcyText = newMovementImpl(-1)
	movementFields.Statistics1 = newMovementImpl(11)
	movementFields.Statistics2 = newMovementImpl(10)
	movementFields.Statistics3 = newMovementImpl(12)
	movementFields.Statistics4 = newMovementImpl(-1)
	movementFields.Description1 = newMovementImpl(9)

	return movementFields
}
