package main

import (
	"flag"
	"fmt"
	lib "kpi_lab_2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	res, err := lib.PostfixToInfix("4 2 - 3 * 5 +")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
