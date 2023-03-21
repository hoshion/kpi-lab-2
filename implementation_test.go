package kpi_lab_2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	res, err := PostfixToInfix("7 3 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "7 + 3", res)
	}

	res, err = PostfixToInfix("4 2 + 1 -")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 + 2 - 1", res)
	}

	res, err = PostfixToInfix("3 6 + 9 2 - *")
	if assert.Nil(t, err) {
		assert.Equal(t, "( 3 + 6 ) * ( 9 - 2 )", res)
	}

	res, err = PostfixToInfix("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "( 4 - 2 ) * 3 + 5", res)
	}

	res, err = PostfixToInfix("7 2 + 4 7 9 + * 9 3 + ^ *")
	if assert.Nil(t, err) {
		assert.Equal(t, "( 7 + 2 ) * ( 4 * ( 7 + 9 ) ) ^ ( 9 + 3 )", res)
	}

	res, err = PostfixToInfix("")
	if err == nil {
		assert.Equal(t, "it's not a number or operator", err)
	}

	res, err = PostfixToInfix("?;\"%\":\"№?")
	if err == nil {
		assert.Equal(t, "it's not a number or operator", err)
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("2 2 +")
	fmt.Println(res)

	// Output:
	// 2 + 2
}
