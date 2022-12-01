package day01

import (
	"aoc/common"
	"strconv"
	"strings"
)

func Solve1(input string) int {
	lines := strings.Split(input, "\n")

	elf := 0
	max := 0
	for _, line := range lines {
		if len(line) > 0 {
			calories, _ := strconv.Atoi(line)
			elf += calories
			if elf > max {
				max = elf
			}
		} else {
			elf = 0
		}
	}

	return max
}

func Part1() {
	name := "Day #01 - part 1"

	common.TestOutput(name+" - input 1", 24000, Solve1(input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day01/input.txt")))
}
