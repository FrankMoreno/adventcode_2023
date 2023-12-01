package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

// ReadInput takes a file name and splits it into lines based on provided separator
func ReadInput(fileName, separator string) ([]string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), separator)

	return lines, nil
}

func ReadInputStream(fileName string) (*bufio.Scanner, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	return scanner, nil
}

func Max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}

	return val2
}
