package kpi_lab_2

import (
	"fmt"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buffer, rerr := io.ReadAll(ch.Input)
	if rerr != nil {
		return rerr
	}

	expression := string(buffer)
	res, err := PostfixToInfix(expression)

	if err != nil {
		return err
	}

	if ch.Output == nil {
		fmt.Println(res)
	} else {
		_, werr := ch.Output.Write([]byte(res))

		if werr != nil {
			return werr
		}
	}

	return nil
}
