package day12

import (
	"aoc/common"
)

func Solve2(input string) int {
	lines, start, _ := parseInput(input, 'E', 'S')

	canMove := func(from, to byte) bool {
		return int(from)-int(to) <= 1
	}

	isEnd := func(c Coords) bool {
		b, _ := valAtCoords(lines, c)
		return b == 'a'
	}

	return bfs(lines, start, isEnd, canMove)
}

func Part2() {
	name := "Day #12 - part 2"

	common.TestOutput(name+" - input 2", 29, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day12/input.txt")))
}
