package main

import (
	"fmt"
)

func intToWord(param int) string {
	word := [21]string{
		"zero", "one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen",
		"sixteen", "seventeen", "eighteen", "nineteen", "twenty",
	}
	if param < 0 {
		return "incorrect number"
	}
	if param > 20 {
		return "big number"
	}
	return word[param]
}

func myCalculation(num1, num2 int, operation string) string {
	result := 0
	switch operation {
	case "+":
		result = num1 + num2
	case "*":
		result = num1 * num2
	default:
		panic("invalid operation")
	}
	return fmt.Sprintf("%s %s %s = %s",
		intToWord(num1),
		operation,
		intToWord(num2),
		intToWord(result))
}
func main() {
	//запускает цикл, от 1 до 10, и создается переменная multiplier
	//для параметра функции myCalculatoin, и выводятся операции
	//сложение и умножение с числами.
	for i := 1; i <= 10; i++ {
		multiplier := 2
		fmt.Println(myCalculation(i, multiplier, "+"))
		fmt.Println("__________________________________")
		fmt.Println(myCalculation(i, multiplier, "*"))

	}
}
