package main

import (
	"fmt"
	"strconv"
	"strings"

	"adventofcode/m/pkg/files"
)

var fileName = "../_testcase/input.txt"

type Game struct {
	id   int
	sets []Set
}

type Set struct {
	red   int
	blue  int
	green int
}

func readingLine(line string) Game {
	// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	colonSplitted := strings.Index(line, ":")

	firstCol := strings.Index(line, " ")
	id := line[firstCol+1 : colonSplitted]

	sets := make([]Set, 0)
	for _, rawSet := range strings.Split(line[colonSplitted+1:], ";") {
		colors := strings.Split(strings.TrimSpace(rawSet), ", ")
		red, blue, green := 0, 0, 0
		for _, part := range colors {
			colorRoll := strings.Split(part, " ")
			val, err := strconv.Atoi(colorRoll[0])
			if err != nil {
				panic(err)
			}
			switch colorRoll[1] {
			case "red":
				red = val
			case "blue":
				blue = val
			case "green":
				green = val
			}
		}
		sets = append(sets, Set{
			red:   red,
			blue:  blue,
			green: green,
		})
	}

	idInt, _ := strconv.Atoi(id)

	return Game{
		id:   idInt,
		sets: sets,
	}
}

func Calculate(filePath string) int {
	lines := files.ReadFiles(filePath)

	games := make([]Game, 0)
	for _, line := range lines {
		games = append(games, readingLine(line))
	}

	minPow := 0
	for _, game := range games {
		minRed, minBlue, minGreen := 0, 0, 0
		for _, set := range game.sets {
			if set.red > minRed {
				minRed = set.red
			}
			if set.blue > minBlue {
				minBlue = set.blue
			}
			if set.green > minGreen {
				minGreen = set.green
			}
		}
		minPow += minRed * minBlue * minGreen
	}

	return minPow
}

func main() {
	filePath := fileName
	minPow := Calculate(filePath)
	fmt.Println(minPow)
}
