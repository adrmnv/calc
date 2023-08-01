package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanDigits = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func isRomanNumber(str string) bool {
	for _, r := range str {
		if _, ok := romanDigits[r]; !ok {
			return false
		}
	}
	return true
}

func romanToArabic(roman string) (int, error) {
	total := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		currentValue := romanDigits[rune(roman[i])]
		if currentValue < prevValue {
			total -= currentValue
		} else {
			total += currentValue
		}
		prevValue = currentValue
	}

	return total, nil
}

func arabicToRoman(arabic int) (string, error) {
	if arabic <= 0 {
		return "", fmt.Errorf("римские числа не могут быть меньше или равны нулю")
	}
	if arabic > 3999 {
		return "", fmt.Errorf("римские числа могут быть только положительными и меньше 4000")
	}

	roman := ""
	for _, digit := range []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	} {
		for arabic >= digit.value {
			roman += digit.symbol
			arabic -= digit.value
		}
	}

	return roman, nil
}

func calculate(a, b interface{}, operator string) (int, error) {
	switch operator {
	case "+":
		return a.(int) + b.(int), nil
	case "-":
		return a.(int) - b.(int), nil
	case "*":
		return a.(int) * b.(int), nil
	case "/":
		if b.(int) == 0 {
			return 0, fmt.Errorf("нельзя делить на 0")
		}
		return a.(int) / b.(int), nil
	default:
		return 0, fmt.Errorf("некорректный оператор")
	}
}

func main() {
	fmt.Println("калькулятор поддерживает арабские и римские числа, от 1 до 10")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)

	elements := strings.Split(input, " ")
	if len(elements) != 3 {
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	aStr, operator, bStr := elements[0], elements[1], elements[2]

	var a, b interface{}

	if a, err = strconv.Atoi(aStr); err != nil {
		a, err = romanToArabic(aStr)
	}
	if err != nil || a.(int) < 1 || a.(int) > 10 {
		panic("ошибка: некорректный операнд a")
	}

	if b, err = strconv.Atoi(bStr); err != nil {
		b, err = romanToArabic(bStr)
	}
	if err != nil || b.(int) < 1 || b.(int) > 10 {
		panic("ошибка: некорректный операнд b")
	}

	if isRomanNumber(aStr) && !isRomanNumber(bStr) || !isRomanNumber(aStr) && isRomanNumber(bStr) {
		fmt.Println("ошибка: используются одновременно разные системы счисления.")
		return
	}

	result, err := calculate(a, b, operator)
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}

	if isRomanNumber(aStr) && isRomanNumber(bStr) {
		romanResult, err := arabicToRoman(result)
		if err != nil {
			fmt.Println("ошибка:", err)
			return
		}
		fmt.Println(romanResult)
	} else {
		fmt.Println(result)
	}
}
