package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var fileName = "DAY01/_testcase/input.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
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
}
