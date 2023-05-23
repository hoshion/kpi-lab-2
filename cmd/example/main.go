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
	var readFile *os.File
	var writeFile *os.File
	var err error

	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else {
		readFile, err = os.Open(*fileExpression)

		if err != nil {
			panic(err)
		}

		reader = io.Reader(readFile)
	}

	if *outputExpression != "" {
		writeFile, err = os.Create(*outputExpression)

		if err != nil {
			panic(err)
		}

		writer = io.Writer(writeFile)
	}

	handler := lib.ComputeHandler{Input: reader, Output: writer}
	err = handler.Compute()

	if err != nil {
		panic(err)
	}

	defer readFile.Close()
	defer writeFile.Close()
}
