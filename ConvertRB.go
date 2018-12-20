package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
	"reflect"
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
			s := reflect.ValueOf(&movementFields).Elem()
			typeOfT := s.Type()
			for i := 0; i < s.NumField(); i++ {
				f := s.Field(i)
				x := genericGetValue(f.Interface().(MovementImpl), rec)
				movementPointer := reflect.ValueOf(&movement)
				s := movementPointer.Elem()
				movementField := s.FieldByName(typeOfT.Field(i).Name)
				movementField.SetString(x)
			}

			accountMovements.AddMovement(movement)
		}
	}

	xmlstring, err := xml.MarshalIndent(accountMovements, "", "    ")

	check(err)
	xmlstring = []byte(xml.Header + string(xmlstring))
	fmt.Fprintf(w, "%s\n", xmlstring)

}

func genericGetValue(movementFieldImpl MovementImpl, rec []string) string {
	v := reflect.ValueOf(movementFieldImpl.function)
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	argv[0] = reflect.ValueOf(rec)
	argv[1] = reflect.ValueOf(movementFieldImpl.index)
	result := v.Call(argv)[0]
	x := result.String()
	return x
}
