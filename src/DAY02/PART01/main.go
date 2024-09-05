package main

import (
	"fmt"
	"strconv"
	"strings"

	"adventofcode/m/pkg/files"
)

var fileName = "../_testcase/input.txt"
var redValue = 12
var greenValue = 13
var blueValue = 14

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

	sumPossibleGameId := 0
	for _, game := range games {
		possibleSets := 0
		for _, set := range game.sets {
			if set.red <= redValue && set.green <= greenValue && set.blue <= blueValue {
				possibleSets++
			}
		}
		if possibleSets == len(game.sets) {
			sumPossibleGameId += game.id
		}
	}

	return sumPossibleGameId
}

func main() {
	filePath := fileName
	sumPossibleGameId := Calculate(filePath)
	fmt.Println(sumPossibleGameId)
}
