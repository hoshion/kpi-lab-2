package main

import (
	"flag"
	"fmt"
	lib "kpi_lab_2"
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

	var expression string

	if *fileExpression == "" {
		expression = *inputExpression
	} else {
		//var r io.Reader
		//r, ferr := os.Open(*fileExpression)
		//
		//if ferr != nil {
		//	panic("Error with file")
		//}

		expression = "4 2 - 3 * 5 +"
	}

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	res, err := lib.PostfixToInfix(expression)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
