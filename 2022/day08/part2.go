package day08

import (
	"aoc/common"
)

func Solve2(input string) int {
	return 1
}

func Part2() {
	name := "Day #08 - part 2"

	common.TestOutput(name+" - input 1", 24933642, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day08/input.txt")))
}
