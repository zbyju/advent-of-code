package day02

import (
	"aoc/common"
	"strings"
)

func Solve2(input string) (score int) {
	lines := strings.Split(input, "\n")

	scores := map[string]int{
		"A X": 3, "A Y": 4, "A Z": 8,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 2, "C Y": 6, "C Z": 7,
	}

	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		score += scores[line]
	}
	return score
}

func Part2() {
	name := "Day #02 - part 2"

	common.TestOutput(name+" - input 1", 12, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day02/input.txt")))
}
