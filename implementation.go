package kpi_lab_2

import (
	"errors"
	"strconv"
	"strings"
)

var (
	operators = "-+*/^"
)

type Node struct {
	value  string
	left   *Node
	right  *Node
	parent *Node
}

func ParseSymbol(strs []string, parent *Node, index *int) *Node {
	node := &Node{
		strs[*index],
		nil,
		nil,
		parent,
	}
	*index--
	if *index < 0 {
		return node
	}
	if strings.ContainsAny(strs[*index], operators) {
		node.right = ParseSymbol(strs, node, index)
	} else {
		node.right = &Node{
			strs[*index],
			nil,
			nil,
			node,
		}
	}
	*index--
	if strings.ContainsAny(strs[*index], operators) {
		node.left = ParseSymbol(strs, node, index)
	} else {
		node.left = &Node{
			strs[*index],
			nil,
			nil,
			node,
		}
	}

	return node
}

func IsBracket(value string, parentValue string) bool {
	if strings.ContainsAny(value, "-+") && strings.ContainsAny(parentValue, "*/^") {
		return true
	}
	if strings.ContainsAny(value, "-+*/") && strings.ContainsAny(parentValue, "^") {
		return true
	}

	return false
}

func ParseNode(node *Node, result *string) {
	if node.parent != nil && IsBracket(node.value, node.parent.value) {
		*result += "( "
	}
	if node.left != nil {
		ParseNode(node.left, result)
	}
	*result += node.value + " "
	if node.right != nil {
		ParseNode(node.right, result)
	}
	if node.parent != nil && IsBracket(node.value, node.parent.value) {
		*result += ") "
	}
}

func PostfixToInfix(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("string is empty")
	}

	symbols := strings.Split(input, " ")

	for _, value := range symbols {
		if len(value) != 1 {
			_, err := strconv.Atoi(value)
			if err != nil {
				return "", errors.New("it's not a number")
			}
		} else if !strings.ContainsAny(value, "0123456789+-*/^") {
			return "", errors.New("it's not a number or operator")
		}
	}

	index := len(symbols) - 1
	point := &index

	root := ParseSymbol(symbols, nil, point)

	result := ""

	ParseNode(root, &result)

	return result, nil
}
