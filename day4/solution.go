package day4

import (
	"strconv"
	"strings"

	"github.com/FrankMoreno/adventcode_2023/collections"
)

func Part1(inputs []string) int {
	total := 0
	for _, input := range inputs {
		winningNumbersString, myNumbersString := separateInput(input)

		winningNumbersSet := collections.NewSet[int]()
		for _, number := range strings.Fields(winningNumbersString) {
			val, _ := strconv.Atoi(number)
			winningNumbersSet.Add(val)
		}

		localTotal := 0

		for _, number := range strings.Split(myNumbersString, " ") {
			val, err := strconv.Atoi(number)
			if err == nil && winningNumbersSet.Contains(val) {
				if localTotal == 0 {
					localTotal = 1
				} else {
					localTotal *= 2
				}
			}
		}

		total += localTotal
	}

	return total
}

func Part2(inputs []string) int {
	mulipliers := map[int]int{}
	total := 0

	for cardNumber, input := range inputs {
		winningNumbersString, myNumbersString := separateInput(input)

		winningNumbersSet := collections.NewSet[int]()
		for _, number := range strings.Fields(winningNumbersString) {
			val, _ := strconv.Atoi(number)
			winningNumbersSet.Add(val)
		}

		winningCards := 0
		for _, number := range strings.Fields(myNumbersString) {
			val, _ := strconv.Atoi(number)
			if winningNumbersSet.Contains(val) {
				winningCards += 1
			}
		}

		mulipliers[cardNumber] += 1
		total += mulipliers[cardNumber]

		for i := cardNumber + 1; i <= cardNumber+winningCards; i++ {
			mulipliers[i] += mulipliers[cardNumber]
		}
	}

	return total
}

func separateInput(input string) (string, string) {
	noCardInput := strings.Split(input, ":")
	splitInput := strings.Split(noCardInput[1], "|")

	return strings.TrimSpace(splitInput[0]), strings.TrimSpace(splitInput[1])
}
