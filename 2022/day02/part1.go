package day02

import (
	"aoc/common"
	"strings"
)

func Solve1(input string) (score int) {
	lines := strings.Split(input, "\n")

	scores := map[string]int{
		"A X": 4, "A Y": 8, "A Z": 3,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 7, "C Y": 2, "C Z": 6,
	}

	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		score += scores[line]
	}
	return score
}

func Part1() {
	name := "Day #02 - part 1"

	common.TestOutput(name+" - input 1", 15, Solve1(Input1))
	common.TestOutput(name+" - input 2", 22, Solve1(Input2))
	common.PrintOutput(name, Solve1(common.Readfile("./day02/input.txt")))
}
