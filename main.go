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
	for i := 1; i <= 20; i++ {
		multiplier := 3
		fmt.Println(myCalculation(i, multiplier, "+"))
		fmt.Println(myCalculation(i, multiplier, "*"))
		fmt.Println("--------------------")
		fmt.Println("Hello world")

		fmt.Println(intToWord(19))
	}
}
