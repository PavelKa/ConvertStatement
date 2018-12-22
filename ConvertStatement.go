package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"os"
)

func main() {

	inputFile := flag.String("i", "", "Input file")
	outputFile := flag.String("o", "", "Output file")
	inputType := flag.String("it", "", "Input type (RBSTATEMENTCSV,FAKTUROID)")
	flag.Parse()

	fmt.Println("inputFile:", *inputFile)
	fmt.Println("outputFile:", *outputFile)
	fmt.Println("inputType:", *inputType)
	fo, err := os.Create(*outputFile)
	check(err)
	defer fo.Close() // make sure it gets closed after

	fi, err := os.Open(*inputFile)
	check(err)
	defer fi.Close()

	switch *inputType {
	case "RBSTATEMENTCSV":
		dat := charmap.Windows1250.NewDecoder().Reader(fi)
		r := csv.NewReader(dat)
		r.Comma = ';'
		target := NewAccountMovements()
		convert(fo, r, rbs, &target)
	case "FAKTUROID":
		//	dat := charmap.Windows1250.NewDecoder().Reader(fi)
		r := csv.NewReader(fi)
		r.Comma = ','
		target := newWinstrom()
		convert(fo, r, fakturoidVydane, &target)
	default:
		fmt.Printf("Unsupported conversion")
	}

}
