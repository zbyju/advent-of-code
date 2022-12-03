package day03

import (
	"aoc/common"
	"strings"
)

func LetterToScore(letter rune) int {
	if letter >= 'a' && letter <= 'z' {
		return int(letter) - int('a') + 1
	} else if letter >= 'A' && letter <= 'Z' {
		return int(letter) - int('A') + 26 + 1
	}
	return 0
}

func Solve1(input string) (sum int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		// First half < mid; Second half >= mid
		mid := len(line) / 2
		chars := make(map[rune]bool)

		for i, c := range line {
			if i < mid {
				chars[c] = true
			} else {
				if chars[c] {
					sum += LetterToScore(c)
					break
				}
			}
		}
	}
	return sum
}

func Part1() {
	name := "Day #03 - part 1"

	common.TestOutput(name+" - input 1", 157, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day03/input.txt")))
}
