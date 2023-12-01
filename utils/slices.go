package utils

import "strconv"

func StringtoIntSlice(stringList []string) ([]int, error) {
	var intSlice = make([]int, len(stringList))

	for i := 0; i < len(stringList); i++ {
		convertedString, err := strconv.Atoi(stringList[i])
		if err != nil {
			return nil, err
		}

		intSlice[i] = convertedString
	}

	return intSlice, nil
}

func Sum[T int | float32 | float64](inputs []T) T {
	var total T

	for _, input := range inputs {
		total += input
	}

	return total
}

func Avg[T int | float32 | float64](inputs []T) T {
	return Sum(inputs) / T(len(inputs))
}
