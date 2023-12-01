package day1

import (
	"unicode"
)

var numberMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func Solution(inputs []string) int {
	accumulator := 0
	for _, input := range inputs {
		accumulator += 10 * getFirstDigit(input)
		accumulator += getLastDigit(input)
	}

	return accumulator
}

func getFirstDigit(input string) int {
	for i := 0; i < len(input); i += 1 {
		char := rune(input[i])
		if unicode.IsDigit(char) {
			return int(char - '0')
		}

		if val := checkForStringNumberForward(input, 0, i+1); val != -1 {
			return val
		}
	}

	return -1
}

func getLastDigit(input string) int {
	for i := len(input) - 1; i >= 0; i -= 1 {
		char := rune(input[i])
		if unicode.IsDigit(char) {
			return int(char - '0')
		}

		if val := checkForStringNumberBackward(input, i, len(input)); val != -1 {
			return val
		}
	}

	return -1
}

func checkForStringNumberForward(input string, start, end int) int {
	for i := end; i >= start; i-- {
		if val, ok := numberMap[input[i:end]]; ok {
			return val
		}
	}

	return -1
}

func checkForStringNumberBackward(input string, start, end int) int {
	for i := start; i <= end; i++ {
		if val, ok := numberMap[input[start:i]]; ok {
			return val
		}
	}

	return -1
}
