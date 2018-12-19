package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	inputFile := flag.String("i", "", "Input file")
	outputFile := flag.String("o", "", "Output file")
	inputType := flag.String("it", "", "Input type (RB)")
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
	convert(fo, fi)
}
