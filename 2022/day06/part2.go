package day06

import (
	"aoc/common"
)

func Solve2(input string) int {
	return Solve1(input, 14)
}

func Part2() {
	name := "Day #06 - part 2"

	common.TestOutput(name+" - input 1", 19, Solve2(Input1))
	common.TestOutput(name+" - input 2", 23, Solve2(Input2))
	common.TestOutput(name+" - input 3", 23, Solve2(Input3))
	common.TestOutput(name+" - input 4", 29, Solve2(Input4))
	common.TestOutput(name+" - input 5", 26, Solve2(Input5))
	common.PrintOutput(name, Solve2(common.Readfile("./day06/input.txt")))
}
