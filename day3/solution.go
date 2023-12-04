package day3

import (
	"fmt"
	"strconv"
	"unicode"
)

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
	if s.validIndex(row+1, column) && s.engine[row+1][column] != '.' && !unicode.IsDigit(rune(s.engine[row+1][column])) {
		return true
	}

	if s.validIndex(row-1, column) && s.engine[row-1][column] != '.' && !unicode.IsDigit(rune(s.engine[row-1][column])) {
		return true
	}

	if s.validIndex(row, column+1) && s.engine[row][column+1] != '.' && !unicode.IsDigit(rune(s.engine[row][column+1])) {
		return true
	}

	if s.validIndex(row, column-1) && s.engine[row][column-1] != '.' && !unicode.IsDigit(rune(s.engine[row][column-1])) {
		return true
	}

	if s.validIndex(row+1, column+1) && s.engine[row+1][column+1] != '.' && !unicode.IsDigit(rune(s.engine[row+1][column+1])) {
		return true
	}

	if s.validIndex(row-1, column-1) && s.engine[row-1][column-1] != '.' && !unicode.IsDigit(rune(s.engine[row-1][column-1])) {
		return true
	}

	if s.validIndex(row+1, column-1) && s.engine[row+1][column-1] != '.' && !unicode.IsDigit(rune(s.engine[row+1][column-1])) {
		return true
	}

	if s.validIndex(row-1, column+1) && s.engine[row-1][column+1] != '.' && !unicode.IsDigit(rune(s.engine[row-1][column+1])) {
		return true
	}

	return false
}

func (s schematic) checkForDigit(row, column int) bool {
	return unicode.IsDigit(rune(s.engine[row][column]))
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
				fmt.Println(number)
				val, _ := strconv.Atoi(number)
				total += val
				number = ""
				valid = false
			}
		}
	}

	return total
}
