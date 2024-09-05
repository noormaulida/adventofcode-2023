package main

import (
	"fmt"
	"strconv"
	"strings"

	"adventofcode/m/pkg/files"
)

var fileName = "../_testcase/input.txt"

func Calculate(filePath string) int {
	lines := files.ReadFiles(fileName)

	total := 0
	for _, line := range lines {
		splittedLine := strings.Split(line, "")

		numbers := []int{}

		for _, letter := range splittedLine {
			if numberFound, err := strconv.Atoi(letter); err == nil {
				numbers = append(numbers, numberFound)
			}
		}

		if len(numbers) > 1 {
			first, last := numbers[0], numbers[len(numbers)-1]
			total += (first * 10) + last
		} else if len(numbers) == 1 {
			num := numbers[0]
			total += (num * 10) + num
		}
	}

	fmt.Println(total)
	return total
}

func main() {
	sumCalibrationValue := Calculate(fileName)
	fmt.Println(sumCalibrationValue)
}
