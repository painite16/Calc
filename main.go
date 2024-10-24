package main

import (
	"errors"
	"unicode"
)

/*
	func Calc(expression string) (float64, error){
		symbols:=make([]string, 0)

		for i := 0; i < len(expression); i++ {
			symbols = append(symbols,string(expression[i]))
		}
		var op1 float64
		var op2 float64
		for _, symb := range symbols{
			if symb != "(" || symb != ")" || symb != "-" || symb != "+" || symb != "*" || symb != "/" {
				op1 = strconv.ParseFloat()
			}
		}
	}
*/
func IsOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}

func GetPriority(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

func InfixToPostfix(expr string) (string, error) {
	var postfix []rune
	var stack []rune

	for _, char := range expr {
		if unicode.IsDigit(char) {
			postfix = append(postfix, char)
		} else if IsOperator(char) {
			for len(stack) > 0 && GetPriority(string(stack[len(stack)-1])) >= GetPriority(string(char)) {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, char)
		} else if char == '(' {
			stack = append(stack, char)
		} else if char == ')' {
			for stack[len(stack)-1] != '(' {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		} else {
			return "", errors.New("Invalid expression")
		}
	}
	for len(stack) > 0 {
		if stack[len(stack)-1] == '(' {
			return "", errors.New("Invalid expression")
		}
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return string(postfix), nil
}

func EvaluatePostfix(expr string) (float64, error) {
	var stack []float64

	for _, char := range expr {
		if unicode.IsDigit(char) {
			stack = append(stack, float64(char-'0'))
		} else if IsOperator(char) {
			if len(stack) < 2 {
				return 0, errors.New("Invalid expression")
			}
			op2 := stack[len(stack)-1]

			stack = stack[:len(stack)-1]
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var result float64

			switch char {
			case '+':
				result = op1 + op2
			case '-':
				result = op1 - op2
			case '*':
				result = op1 * op2
			case '/':
				if op2 == 0 {
					return 0, errors.New("Division by zero")
				}
				result = op1 / op2
			}
			stack = append(stack, result)
		} else {
			return 0, errors.New("Invalid expression")
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("Invalid expression")
	}
	return stack[0], nil
}

func Calc(expression string) (float64, error) {
	postfixExpr, err := InfixToPostfix(expression)
	if err != nil {
		return 0, err
	}

	result, err := EvaluatePostfix(postfixExpr)
	if err != nil {
		return 0, err
	}
	return result, nil
}
