package day2

import (
	"strconv"
	"strings"
)

type game struct {
	number int
	blue   int
	red    int
	green  int
}

func Part1(inputs []string) int {
	total := 0
	for idx, input := range inputs {
		gameInfo := strings.Split(input, ":")
		if parseGame(gameInfo[1], idx+1) {
			total += idx + 1
		}
	}

	return total
}

func Part2(inputs []string) int {
	total := 0
	for idx, input := range inputs {
		gameInfo := strings.Split(input, ":")
		parsedGame := parseRound(gameInfo[1], idx+1)
		total += parsedGame.blue * parsedGame.red * parsedGame.green
	}

	return total
}

func parseGame(input string, number int) bool {
	rounds := strings.Split(input, ";")
	for _, round := range rounds {
		parsedGame := game{}
		parsedGame.number = number
		dice := strings.Split(round, ",")
		for _, die := range dice {
			die = strings.TrimSpace(die)
			info := strings.Split(die, " ")
			amount, _ := strconv.Atoi(info[0])
			if info[1] == "blue" {
				parsedGame.blue += amount
			}

			if info[1] == "green" {
				parsedGame.green += amount
			}

			if info[1] == "red" {
				parsedGame.red += amount
			}
		}
		if !parsedGame.feasible() {
			return false
		}
	}

	return true
}

func parseRound(input string, number int) game {
	parsedGame := game{}
	parsedGame.number = number
	rounds := strings.Split(input, ";")
	for _, round := range rounds {
		dice := strings.Split(round, ",")
		for _, die := range dice {
			die = strings.TrimSpace(die)
			info := strings.Split(die, " ")
			amount, _ := strconv.Atoi(info[0])
			if info[1] == "blue" {
				if amount > parsedGame.blue {
					parsedGame.blue = amount
				}
			}

			if info[1] == "green" {
				if amount > parsedGame.green {
					parsedGame.green = amount
				}
			}

			if info[1] == "red" {
				if amount > parsedGame.red {
					parsedGame.red = amount
				}
			}
		}
	}

	return parsedGame
}

func (g game) feasible() bool {
	return g.red <= 12 && g.green <= 13 && g.blue <= 14
}
