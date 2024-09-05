package main

import (
	"fmt"
	"strings"
	"unicode"

	"adventofcode/m/pkg/files"
)

var fileName = "../_testcase/input.txt"

var numberInWords = map[string]int{
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

func isContainSpelledNumber(line string) (bool, int) {
	for k, v := range numberInWords {
		if strings.Contains(line, k) {
			return true, v
		}
	}
	return false, 0
}

func getFirstNumber(line string) int {
	for i := 0; i < len(line); i++ {
		if found, d := isContainSpelledNumber(line[:i]); found {
			return d
		} else if unicode.IsDigit(rune(line[i])) {
			return int(line[i] - '0')
		}
	}
	panic("Error " + line)
}

func getLastNumber(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if found, d := isContainSpelledNumber(line[i:]); found {
			return d
		} else if unicode.IsDigit(rune(line[i])) {
			return int(line[i] - '0')
		}
	}
	panic("Error " + line)
}

func Calculate(filePath string) int {
	lines := files.ReadFiles(fileName)

	total := 0
	for _, line := range lines {
		first := getFirstNumber(line)
		last := getLastNumber(line)

		total += (first * 10) + last
	}
	return total
}

func main() {
	filePath := fileName
	sumCalibrationValue := Calculate(filePath)
	fmt.Println(sumCalibrationValue)
}
