package kpi_lab_2

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func TestComputeHandler_Compute(t *testing.T) {
	var reader io.Reader = strings.NewReader("5 5 ?")
	handler := ComputeHandler{Input: reader}
	err := handler.Compute()
	assert.NotEqual(t, nil, err)

	filePath := "result.txt"
	reader = strings.NewReader("5 5 -")
	wfile, _ := os.Create(filePath)
	writer := io.Writer(wfile)
	handler = ComputeHandler{reader, writer}
	err = handler.Compute()
	if assert.Nil(t, err) {
		rfile, _ := os.Open(filePath)
		reader = io.Reader(rfile)
		buffer, _ := io.ReadAll(reader)
		expression := string(buffer)
		assert.Equal(t, "5 - 5", expression)
	}
}
