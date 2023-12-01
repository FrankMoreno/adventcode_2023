package utils

import (
	"fmt"
	"time"
)

func Benchmark(input []string, wrappedFunc func([]string) int, iterations int) {
	start := time.Now()
	solution := 0
	for i := 0; i < iterations; i++ {
		solution = wrappedFunc(input)
	}
	end := time.Since(start)

	fmt.Printf("Solution: %d\n", solution)
	fmt.Printf("Runtime: %d microseconds\n", end.Microseconds())
}

func BenchmarkAverage(input []string, wrappedFunc func([]string) int, iterations int) {
	runTimes := []int{}
	solution := 0
	for i := 0; i < iterations; i++ {
		start := time.Now()
		solution = wrappedFunc(input)
		end := time.Since(start)
		runTimes = append(runTimes, int(end.Microseconds()))
	}

	fmt.Printf("Solution: %d\n", solution)
	fmt.Printf("Runtime: %d microsends\n", Avg(runTimes))
}
