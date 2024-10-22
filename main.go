package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

const (
	ErrorInExpression = "There is an error in the expression"
	ErrorInBrackets   = "There is an error in the brackets"
	DivideByZero      = "Division by zero"
)

func Calc(expression string) (float64, error) {
	if err := validateExpression(expression); err != nil {
		return 0, err
	}
	if err := validateBrackets(expression); err != nil {
		return 0, err
	}

	var numbers []float64
	var operators []rune
	var currentNum string

	for _, char := range expression {
		if unicode.IsDigit(char) || char == '.' {
			currentNum += string(char)
		} else {
			if currentNum != "" {
				num, err := strconv.ParseFloat(currentNum, 64)
				if err != nil {
					return 0, errors.New(ErrorInExpression)
				}
				numbers = append(numbers, num)
				currentNum = ""
			}

			if char == '+' || char == '-' || char == '*' || char == '/' {
				operators = append(operators, char)
			} else if char == '(' {
				operators = append(operators, char)
			} else if char == ')' {
				for len(operators) > 0 && operators[len(operators)-1] != '(' {
					numbers, operators = applyOperator(numbers, operators)
				}
				if len(operators) == 0 {
					return 0, errors.New(ErrorInBrackets)
				}
				operators = operators[:len(operators)-1]
			}
		}
	}

	for i := 0; i < len(expression)-1; i++ {
		if (expression[i] == '+' || expression[i] == '-' || expression[i] == '*' || expression[i] == '/') &&
			(expression[i+1] == '+' || expression[i+1] == '-' || expression[i+1] == '*' || expression[i+1] == '/') {
			return 0, errors.New(ErrorInExpression)
		}
	}

	if currentNum != "" {
		num, err := strconv.ParseFloat(currentNum, 64)
		if err != nil {
			return 0, errors.New(ErrorInExpression)
		}
		numbers = append(numbers, num)
	}

	for len(operators) > 0 {
		numbers, operators = applyOperator(numbers, operators)
	}

	if len(numbers) != 1 {
		return 0, errors.New(ErrorInExpression)
	}

	return numbers[0], nil
}

func validateExpression(expression string) error {
	if len(expression) == 0 {
		return errors.New(ErrorInExpression)
	}

	if unicode.IsSymbol(rune(expression[0])) || unicode.IsSymbol(rune(expression[len(expression)-1])) {
		return errors.New(ErrorInExpression)
	}

	for j := 1; j < len(expression); j++ {
		if unicode.IsSymbol(rune(expression[j])) && unicode.IsSymbol(rune(expression[j-1])) {
			return errors.New(ErrorInExpression)
		}
	}

	if expression[len(expression)-1] == '+' || expression[len(expression)-1] == '-' || expression[len(expression)-1] == '*' || expression[len(expression)-1] == '/' {
		return errors.New(ErrorInExpression)
	}

	return nil
}

func validateBrackets(expression string) error {
	var stack []rune

	for _, char := range expression {
		if char == '(' {
			stack = append(stack, char)
		} else if char == ')' {
			if len(stack) == 0 {
				return errors.New(ErrorInBrackets)
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return errors.New(ErrorInBrackets)
	}

	return nil
}

func applyOperator(numbers []float64, operators []rune) ([]float64, []rune) {
	if len(numbers) < 2 || len(operators) == 0 {
		return numbers, operators
	}

	num2 := numbers[len(numbers)-1]
	num1 := numbers[len(numbers)-2]
	operator := operators[len(operators)-1]

	numbers = numbers[:len(numbers)-2]
	operators = operators[:len(operators)-1]

	var result float64
	switch operator {
	case '+':
		result = num1 + num2
	case '-':
		result = num1 - num2
	case '*':
		result = num1 * num2
	case '/':
		if num2 == 0 {
			return numbers, operators
		}
		result = num1 / num2
	}

	numbers = append(numbers, result)
	return numbers, operators
}

func main() {
	var expression string
	fmt.Scan(&expression)

	result, err := Calc(expression)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
