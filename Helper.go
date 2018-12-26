package main

import (
	"reflect"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
func getDateF(rec []string, i int) string {

	return rec[i] + "+01:00"

}
func getTypDokl(rec []string, i int) string {

	return "code:FAKTURA"

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
	nf = strings.Replace(nf, "-", "", -1)

	return nf

}

func getAmountF(rec []string, i int) string {

	nf := strings.Replace(rec[i], " ", "", -1)
	nf = strings.Replace(nf, "-", "", -1)
	nf = strings.Replace(nf, ",", ".", -1)

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
			return "D"
		} else {
			return "C"
		}
	} else {
		return "U"
	}
}

func setField(v interface{}, name string, value string) error {
	stype := reflect.ValueOf(v).Elem()
	field := stype.FieldByName(name)
	field.SetString(value)

	return nil
}
