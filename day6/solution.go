package day6

import (
	"math"
	"strconv"
	"strings"

	"github.com/FrankMoreno/adventcode_2023/utils"
)

func Part1(inputs []string) int {
	total := 1
	times, distances := parseInput(inputs)

	for i := 0; i < len(times); i++ {
		total *= quadraticFormula(-1, float64(times[i]), float64(-distances[i]))
	}

	return total
}

func Part2(inputs []string) int {
	totalTime, recordDistance := parseCombinedInput(inputs)

	return quadraticFormula(-1, float64(totalTime), float64(-recordDistance))
}

func parseInput(inputs []string) ([]int, []int) {
	timeStrings := strings.Fields(strings.Split(inputs[0], ":")[1])
	distanceStrings := strings.Fields(strings.Split(inputs[1], ":")[1])

	times, _ := utils.StringtoIntSlice(timeStrings)
	distances, _ := utils.StringtoIntSlice(distanceStrings)

	return times, distances
}

func parseCombinedInput(inputs []string) (int, int) {
	timeStrings := strings.Fields(strings.Split(inputs[0], ":")[1])
	distanceStrings := strings.Fields(strings.Split(inputs[1], ":")[1])

	time, _ := strconv.Atoi(strings.Join(timeStrings, ""))
	dist, _ := strconv.Atoi(strings.Join(distanceStrings, ""))

	return time, dist
}

func quadraticFormula(a, b, c float64) int {
	d := math.Sqrt(math.Pow(b, 2) - (4 * a * c))
	e := 2 * a
	x1 := (-b + d) / e
	x2 := (-b - d) / e

	return int(math.Ceil(x2) - math.Floor(x1) - 1)
}
