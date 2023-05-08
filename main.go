package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"utils"
)

func main() {
	var result string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение для расчёта:")
	input, _ := reader.ReadString('\n')

	splittedInput, err := utils.SplitString(input)
	if err != nil {
		utils.ErrorHandler(err)
	}

	a := splittedInput[0]
	operator := splittedInput[1]
	b := splittedInput[2]

	aValue, bValue, isArabic, err := utils.PrepareOperands(a, b, operator)
	if err != nil {
		utils.ErrorHandler(err)
	}

	resultInt, err := utils.Calculte(aValue, bValue, operator)
	if err != nil {
		utils.ErrorHandler(err)
	}

	if isArabic {
		result = strconv.Itoa(resultInt)
	} else {
		result, err = utils.ArabicToRoman(resultInt)
	}

	if err != nil {
		utils.ErrorHandler(err)
	}
	fmt.Println(result)
}
