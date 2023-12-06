package day3

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/FrankMoreno/adventcode_2023/collections"
)

var directions = [][]int{{1, 0}, {0, 1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}, {-1, 0}, {0, -1}}

type schematic struct {
	engine []string
}

func (s schematic) validIndex(row, column int) bool {
	if row >= len(s.engine) {
		return false
	}

	if row < 0 {
		return false
	}

	if column < 0 {
		return false
	}

	if column >= len(s.engine[row]) {
		return false
	}

	return true
}

func (s schematic) checkForSymbol(row, column int) bool {
	for _, direction := range directions {
		if s.validIndex(row+direction[0], column+direction[1]) && s.engine[row+direction[0]][column+direction[1]] != '.' && !unicode.IsDigit(rune(s.engine[row+direction[0]][column+direction[1]])) {
			return true
		}
	}

	return false
}

func (s schematic) checkForAdjacentNumber(row, column int) [][]int {
	total := [][]int{}
	for _, direction := range directions {
		if s.validIndex(row+direction[0], column+direction[1]) && s.checkForDigit(row+direction[0], column+direction[1]) {
			total = append(total, []int{row + direction[0], column + direction[1]})
		}
	}

	return total
}

func (s schematic) checkForDigit(row, column int) bool {
	return unicode.IsDigit(rune(s.engine[row][column]))
}

func (s schematic) checkForGear(row, column int) bool {
	return s.engine[row][column] == '*'
}

func (s schematic) getFullDigit(row, column int) int {
	full := ""
	for i := column; s.validIndex(row, i) && s.checkForDigit(row, i); i -= 1 {
		full = string(s.engine[row][i]) + full
	}

	for i := column + 1; s.validIndex(row, i) && s.checkForDigit(row, i); i += 1 {
		full = full + string(s.engine[row][i])
	}

	val, _ := strconv.Atoi(full)
	return val
}

func Part1(input []string) int {
	schem := schematic{
		engine: input,
	}
	total := 0
	for x, row := range input {
		number := ""
		valid := false
		for y, _ := range row {
			if schem.checkForDigit(x, y) {
				number += string(schem.engine[x][y])
				if !valid {
					valid = schem.checkForSymbol(x, y)
				}
			} else {
				if valid {
					// fmt.Println(number)
					val, err := strconv.Atoi(number)
					if err != nil {
						fmt.Println("Failed to convert!!")
					}
					total += val
				}
				number = ""
				valid = false
			}
		}

		if valid {
			val, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println("Failed to convert!!")
			}
			total += val
		}
	}

	return total
}

func Part2(input []string) int {
	schem := schematic{
		engine: input,
	}
	total := 0
	for x, row := range input {
		for y, _ := range row {
			if schem.checkForGear(x, y) {
				localTotal := 1
				adjacentNums := schem.checkForAdjacentNumber(x, y)
				unique := collections.NewSet[int]()

				for _, adjacantNum := range adjacentNums {
					fullFigit := schem.getFullDigit(adjacantNum[0], adjacantNum[1])
					if !unique.Contains(fullFigit) {
						unique.Add(fullFigit)
						localTotal *= fullFigit
					}
				}

				if unique.Len() > 1 {
					total += localTotal
				}
			}
		}
	}

	return total
}
