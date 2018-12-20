package main

type MovementField int

type MovementFields struct {
	Amount           MovementField
	Direction        MovementField
	PostingDate      MovementField
	ItemNo           MovementField
	PartnerAccNo     MovementField
	PartnerAccBank   MovementField
	PartnerAccName   MovementField
	ValueDate        MovementField
	ChargesAmount    MovementField
	ChargesCcy       MovementField
	MovementTypeText MovementField
	BankCode         MovementField
	BankCountryID    MovementField
	BankName         MovementField
	AccNoID          MovementField
	AccNoCC          MovementField
	AccName          MovementField
	AccCcy           MovementField
	AccCcyText       MovementField
	Statistics1      MovementField
	Statistics2      MovementField
	Statistics3      MovementField
	Statistics4      MovementField
	Description1     MovementField
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
func newMovementFieldRsB() MovementFields {
	movementFields := MovementFields{}
	movementFields.Amount = 13
	movementFields.Direction = -1
	movementFields.PostingDate = 0
	movementFields.ItemNo = 18
	movementFields.PartnerAccNo = 5
	movementFields.PartnerAccBank = -1
	movementFields.PartnerAccName = 6
	movementFields.ValueDate = 1
	movementFields.ChargesAmount = 17
	movementFields.ChargesCcy = 17
	movementFields.MovementTypeText = 7
	movementFields.BankCode = -1
	movementFields.BankCountryID = -1
	movementFields.BankName = -1
	movementFields.AccNoID = 2
	movementFields.AccNoCC = -1
	movementFields.AccName = 3
	movementFields.AccCcy = -1
	movementFields.AccCcyText = -1
	movementFields.Statistics1 = 10
	movementFields.Statistics2 = 11
	movementFields.Statistics3 = 12
	movementFields.Statistics4 = -1
	movementFields.Description1 = 9
	return movementFields
}
