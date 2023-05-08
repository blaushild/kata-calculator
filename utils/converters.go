package utils

import (
	"fmt"
	"strconv"
)

func RomanToArabic(romanNumber string) (int, error) {
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	previousValue := 0

	for _, c := range romanNumber {
		value, exists := romanValues[c]
		if !exists {
			return 0, fmt.Errorf("Некорректный символ: %c", c)
		}

		if value > previousValue {
			result += value - 2*previousValue
		} else {
			result += value
		}

		previousValue = value
	}

	return result, nil
}

func ArabicToRoman(number int) (string, error) {
	if number < 1 || number > 3999 {
		return "", fmt.Errorf("Недопустимое значение числа: %d. Конвертация в римские числа возможна только для чисел в диапазоне 1..3999.", number)
	}

	romanNumerals := []struct {
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
	}

	romanNum := ""
	for _, numeral := range romanNumerals {
		for number >= numeral.value {
			romanNum += numeral.symbol
			number -= numeral.value
		}
	}

	return romanNum, nil
}

func ConvertOperand(operand string) (int, bool, error) {
	/* преобразует число из string в int
	возвращает значение и флаг
	флаг true -- входное число арабское
	флаг false -- входное число римское
	возвращает ошибку, если не может конвертировать
	или если число за пределами диапазона 1..10 */
	isArabic := true
	value, err := strconv.Atoi(operand)
	if err != nil {
		value, err = RomanToArabic(operand)
		isArabic = false
	}

	if err != nil {
		// if can't convert
		value64, errTemp := strconv.ParseFloat(operand, 64)
		err = errTemp
		if err != nil {
			err = fmt.Errorf("Не могу сконвертировать операнд '%s' в число.", operand)
		} else {
			err = fmt.Errorf("Калькулятор умеет работать только с целыми числами. '%f' -- дробное.\n", value64)
		}
		return 0, false, err
	}

	if 1 > value || value > 10 {
		err = fmt.Errorf("Число %d находится за пределами диапазона 1..10", value)
		return 0, false, err
	}

	return value, isArabic, nil
}
