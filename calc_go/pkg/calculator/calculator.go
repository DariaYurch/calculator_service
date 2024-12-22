package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func brokeString(str string) []string{
	var result []string
	var partOfString strings.Builder
	str = strings.ReplaceAll(str, " ", "")
	for _, elem := range str{
		if elem == '+' || elem == '-' || elem == '/' || elem == '*' || elem == '(' || elem == ')'{
			if partOfString.Len() > 0{
				result = append(result, partOfString.String())
			}
			partOfString.Reset()
			result = append(result, string(elem))
		}else{
			partOfString.WriteRune(elem)
		}

	}
	if partOfString.Len() > 0{
		result = append(result, partOfString.String())
	}
	return result
}

func isNumber(element string) bool{
	_, err := strconv.ParseFloat(element, 64)
	return err == nil
}

func isOperator(element string) bool{
	if element == "+" || element == "-" || element == "/" || element == "*" {
		return true
	}
	return false
}

func getPriority(operator string) int{
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}



func second_func(main_expression []string) ([]string, error){
	var result []string
	var stackOperators []string
	for _, element := range main_expression{
		if isNumber(element){
			result = append(result, element)
		}else if element == "(" {
			stackOperators = append(stackOperators, element)
		}else if element == ")"{
			for len(stackOperators) > 0 && stackOperators[len(stackOperators) - 1] != "("{
				result = append(result, stackOperators[len(stackOperators) - 1])
				stackOperators = stackOperators[:len(stackOperators) - 1]
			}
			if len(stackOperators) == 0{
				return nil, errors.New("Error")
			}
			stackOperators = stackOperators[:len(stackOperators) - 1]
		}else if isOperator(element){
			for len(stackOperators) > 0 && getPriority(stackOperators[len(stackOperators) - 1]) >= getPriority(element){
				result = append(result, stackOperators[len(stackOperators) - 1])
				stackOperators = stackOperators[:len(stackOperators) - 1]
			}
			stackOperators = append(stackOperators, element)
		}else{
			return nil, errors.New("Error")
		}
	}
	for len(stackOperators) > 0{
		if stackOperators[len(stackOperators) - 1] == "("{
			return nil, errors.New("Error")
		}
		result = append(result, stackOperators[len(stackOperators) - 1])
		stackOperators = stackOperators[:len(stackOperators) - 1]
	}
	return result, nil
}

func final_func(workExpression[]string) (float64, error){
	var result []float64

	for _, elem := range workExpression{
		if isNumber(elem){
			num, _ := strconv.ParseFloat(elem, 64)
			result = append(result, num)
		}else if isOperator(elem){
			if len(result) < 2{
				return 0, errors.New("Calculation error. Wrong input")
			}
			a := result[len(result) - 2]
			b := result[len(result) - 1]
			result = result[:len(result) - 2]
			switch elem {
			case "+":
				result = append(result, a + b)
			case "-":
				result = append(result, a - b)
			case "*":
				result = append(result, a * b)
			case "/":
				if b == 0 {
					return 0, errors.New("Calculation error. It can`t be divided by 0")
				}
				result = append(result, a / b)
			default:
				return 0, fmt.Errorf("Error. Unknown operator: %s", elem)
			}
		} else {
			return 0, fmt.Errorf("Error. Invalid token: %s", elem)
		}
	}

	if len(result) != 1 {
		return 0, errors.New("Calculation error. Invalid expression")
	}

	return result[0], nil
}

func Calc(expression string) (float64, error){
	var main_expression []string
	str := expression
	main_expression = brokeString(str)
	result, err := second_func(main_expression)
	r, err := final_func(result)
	return r, err
}
