package utils

import (
	"fmt"
	"os"
	"strings"
)

func ErrorHandler(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func SplitString(input string) ([]string, error) {
	// разбивает полученную строку, используя пробел в качестве разделителя
	// возвращает ошибку, если количество полученных элементов не равно 3
	input = strings.TrimSpace(input)
	splittedInput := strings.Split(input, " ")
	if len(splittedInput) < 3 {
		return splittedInput, fmt.Errorf(ErrorMessages[2])
	}
	if len(splittedInput) > 3 {
		return splittedInput, fmt.Errorf(ErrorMessages[3])
	}
	return splittedInput, nil
}

func PrepareOperands(a, b, operator string) (int, int, bool, error) {
	/* проверяет принадлежность чисел к одной системе записи
	проверят на отрицательный результат для римских чисел
	*/
	valueA, aIsArabic, err := ConvertOperand(a)
	if err != nil {
		return 0, 0, false, err
	}
	valueB, bIsArabic, err := ConvertOperand(b)
	if err != nil {
		return 0, 0, false, err
	}

	if aIsArabic != bIsArabic {
		err = fmt.Errorf(ErrorMessages[1])
	}

	if !(aIsArabic && bIsArabic) && (valueA < valueB) && operator == "-" {
		err = fmt.Errorf(ErrorMessages[0])
	}
	if err != nil {
		return 0, 0, false, err
	}

	return valueA, valueB, aIsArabic, nil
}

func Calculte(a, b int, operator string) (int, error) {
	// считает результат
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return int(a / b), nil
	default:
		return 0, fmt.Errorf("Оператор '%s' не определён.", operator)
	}
}
