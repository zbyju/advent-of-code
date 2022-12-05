package day05

import (
	"aoc/common"
	"strings"
)

func Solve2(input string) (res string) {
	parts := strings.Split(input, "\n\n")
	linesStacks := strings.Split(parts[0], "\n")
	linesMoves := strings.Split(parts[1], "\n")
	stacks := loadCrates(linesStacks)
	moves := parseMoves(linesMoves)

	for _, move := range moves {
		moving := stacks[move.from][len(stacks[move.from])-move.amount:]

		for _, crate := range moving {
			stacks[move.from] = stacks[move.from][:len(stacks[move.from])-1]
			stacks[move.to] = append(stacks[move.to], crate)
		}
	}

	for _, stack := range stacks {
		res += string(stack[len(stack)-1])
	}

	return res
}

func Part2() {
	name := "Day #05 - part 2"

	common.TestOutputStr(name+" - input 1", "MCD", Solve2(Input1))
	common.PrintOutputStr(name, Solve2(common.Readfile("./day05/input.txt")))
}
