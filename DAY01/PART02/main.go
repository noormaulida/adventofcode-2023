package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

var fileName = "DAY01/_testcase/input.txt"

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

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
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

func main() {
	fmt.Println("Opening input file ")
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	if err := sc.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
		return
	}

	total := 0
	for sc.Scan() {
		line := sc.Text()

		first := getFirstNumber(line)
		last := getLastNumber(line)

		total += (first * 10) + last
	}
	fmt.Println(total)
}
