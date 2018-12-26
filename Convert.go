package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
)

func convert(w io.Writer, r *csv.Reader, convDefinition ConversionFields, items interface{}) {

	records, err := r.ReadAll()
	check(err)
	i := 0
	for _, rec := range records {
		i++
		if i > 1 {

			item := items.(ItemList).CreateItem()

			for _, f := range convDefinition {
				x := genericGetValue(f, rec)
				ee := setField(item, f.fieldName, x)
				check(ee)
			}
			items.(ItemList).AddItem(item)

		}
	}

	xmlstring, err := xml.MarshalIndent(items, "", "    ")
	check(err)
	xmlstring = []byte(xml.Header + string(xmlstring))
	fmt.Fprintf(w, "%s\n", xmlstring)

}

func genericGetValue(conversionImpl ConversionImpl, rec []string) string {
	v := reflect.ValueOf(conversionImpl.function)
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	argv[0] = reflect.ValueOf(rec)
	argv[1] = reflect.ValueOf(conversionImpl.index)
	result := v.Call(argv)[0]
	x := result.String()
	return x
}
