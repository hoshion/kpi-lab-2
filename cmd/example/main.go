package main

import (
	"flag"
	"io"
	lib "kpi_lab_2"
	"os"
	"strings"
)

var (
	inputExpression  = flag.String("e", "", "Expression to compute")
	fileExpression   = flag.String("f", "", "File to read")
	outputExpression = flag.String("o", "", "File to write")
)

func main() {
	flag.Parse()

	if (*inputExpression != "" && *fileExpression != "") || (*inputExpression == "" && *fileExpression == "") {
		panic("Error with parameters, should only be one of two - file or expression")
	}

	var reader io.Reader
	var writer io.Writer = nil

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else {
		file, ferr := os.Open(*fileExpression)

		if ferr != nil {
			panic(ferr)
		}

		reader = io.Reader(file)
	}

	if *outputExpression != "" {
		file, ferr := os.Create(*outputExpression)

		if ferr != nil {
			panic(ferr)
		}

		writer = io.Writer(file)
	}

	handler := lib.ComputeHandler{Input: reader, Output: writer}
	err := handler.Compute()

	if err != nil {
		panic(err)
	}
}
