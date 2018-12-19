package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
)

func convert(w io.Writer, f io.Reader) {
	dat := charmap.ISO8859_2.NewDecoder().Reader(f)
	r := csv.NewReader(dat)
	r.Comma = ';'
	accountMovements := NewAccountMovements()
	accountMovements.official = "N"
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	for _, rec := range records {
		if true {
			i = i + 1

			movement := NewMovement()
			movement.PostingDate = rec[0] //Datum provedení
			movement.ValueDate = rec[1]   //    Datum zaúčtování
			movement.AccNoID = rec[2]     //Číslo účtu
			movement.AccName = rec[3]     //Název účtu
			//movement. = rec[0] //Kategorie transakce
			movement.PartnerAccNo = rec[5]     //Číslo protiúčtu
			movement.PartnerAccName = rec[6]   //Název protiúčtu
			movement.MovementTypeText = rec[7] //Typ transakce
			movement.Description1 = rec[8]     //Zpráva
			//movement.ValueDate = rec[8] //Poznámka
			movement.Statistics2 = fmt.Sprintf("%02d", rec[10]) //VS
			movement.Statistics1 = rec[11]                      //KS
			movement.Statistics3 = rec[12]                      //SS
			movement.Amount = rec[13]                           //Zaúčtovaná částka
			movement.AccCcy = rec[14]                           //Měna účtu
			//movement. = rec[14] //Původní částka a měna
			//movement.ValueDate = rec[15] //Původní částka a měna
			movement.ChargesAmount = rec[16] //Poplatek
			movement.ItemNo = rec[18]        //Id transakce

			accountMovements.AddMovement(movement)
		}
	}

	xmlstring, err := xml.MarshalIndent(accountMovements, "", "    ")

	check(err)
	xmlstring = []byte(xml.Header + string(xmlstring))
	fmt.Print(xmlstring)
	fmt.Fprintf(w, "%s\n", xmlstring)

}
