package day12

import (
	"aoc/common"
	"strings"
)

func Solve2(input string) int {
	lines := strings.Split(input, "\n")
	min := -1
	for i := range lines {
		for j := range lines[i] {
			c := Coords{i, j}
			if v, _ := valAtCoords(lines, c); v == 'a' || v == 'S' {
				len := bfs(lines, c)
				if min == -1 || (len != -1 && len < min) {
					min = len
				}
			}
		}
	}
	return min
}

func Part2() {
	name := "Day #12 - part 2"

	common.TestOutput(name+" - input 2", 29, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day12/input.txt")))
}
