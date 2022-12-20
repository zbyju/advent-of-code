package day19

import (
	"aoc/common"
	"strings"
)

func Solve2(input string) int {
	total := 1
	fs := parseInput(strings.Split(input, "\n")[:3])
	for _, f := range fs {
		m := run(f, 32, &map[Key]int{})
		total *= m
	}
	return total
}

func Part2() {
	name := "Day #19 - part 2"

	common.PrintOutput(name, Solve2(common.Readfile("./day19/input.txt")))
}
