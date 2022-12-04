package day04

import (
	"aoc/common"
	"strings"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func doesOverlap(i1, i2 Interval) bool {
	return max(i1.start, i2.start) <= min(i1.end, i2.end)
}

func Solve2(input string) (count int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		sections := ParseLine(line)
		if doesOverlap(sections[0], sections[1]) {
			count++
		}
	}
	return count
}

func Part2() {
	name := "Day #04 - part 2"

	common.TestOutput(name+" - input 1", 4, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day04/input.txt")))
}
