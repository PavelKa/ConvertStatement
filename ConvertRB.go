package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
)

func convert(w io.Writer, f io.Reader, movementFields MovementFields) {
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
		i++
		if i > 1 {
			movement := NewMovement()
			movement.PostingDate = rec[movementFields.PostingDate]
			movement.ValueDate = rec[movementFields.ValueDate]
			movement.AccNoID = rec[movementFields.AccNoID]
			movement.AccName = rec[movementFields.AccName]
			//movement. = rec[0] //Kategorie transakce
			movement.PartnerAccNo = rec[movementFields.PartnerAccNo]
			movement.PartnerAccName = rec[movementFields.PartnerAccName]
			movement.MovementTypeText = rec[movementFields.MovementTypeText]
			movement.Description1 = rec[movementFields.Description1]
			//movement.ValueDate = rec[8] //Poznámka
			movement.Statistics2 = rec[movementFields.Statistics2]
			movement.Statistics1 = rec[movementFields.Statistics1]
			movement.Statistics3 = rec[movementFields.Statistics3]
			movement.Amount = rec[movementFields.Amount]
			//			movement.AccCcy = rec[movementFields.AccCcy]
			//movement. = rec[movementFields.PostingDate]
			//movement.ValueDate = rec[movementFields.PostingDate]
			movement.ChargesAmount = rec[movementFields.ChargesAmount]
			movement.ItemNo = rec[movementFields.ItemNo]

			accountMovements.AddMovement(movement)
		}
	}

	xmlstring, err := xml.MarshalIndent(accountMovements, "", "    ")

	check(err)
	xmlstring = []byte(xml.Header + string(xmlstring))
	fmt.Print(xmlstring)
	fmt.Fprintf(w, "%s\n", xmlstring)

}
