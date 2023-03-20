package kpi_lab_2

import (
	"errors"
	"strconv"
	"strings"
)

var (
	operators = "-+*/^"
)

// Node is an element of an oriented tree (the structure used for representing mathematical formula in different ways)
type Node struct {
	value  string
	left   *Node
	right  *Node
	parent *Node
}

// ParseSymbol creates oriented tree structure, that contains mathematical formula
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

// IsBracket checks weather there should be a bracket in the infix variant or not
func IsBracket(value string, parentValue string) bool {
	if strings.ContainsAny(value, "-+") && strings.ContainsAny(parentValue, "*/^") {
		return true
	}
	if strings.ContainsAny(value, "-+*/") && strings.ContainsAny(parentValue, "^") {
		return true
	}

	return false
}

// ParseNode creates infix variant of formula
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

// PostfixToInfix transforms postfix mathematical formula into natural infix variant.
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

	return strings.TrimSpace(result), nil
}
