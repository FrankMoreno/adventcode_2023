package main

import (
	"log"
	"os"

	"github.com/FrankMoreno/adventcode_2023/day3"
	"github.com/FrankMoreno/adventcode_2023/utils"
)

func main() {
	input, err := utils.ReadInput(os.Args[1], "\n")
	if err != nil {
		log.Panicf("Unable to read file contents: %v", err)
	}

	// fmt.Println(day1.Solution(input))

	utils.BenchmarkAverage(input, day3.Part2, 1)
}
