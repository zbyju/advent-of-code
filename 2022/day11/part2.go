package day11

import (
	"aoc/common"
	"strings"
)

func Solve2(input string) int64 {
	monkeys := []Monkey{}
	for _, monkeyStr := range strings.Split(input, "\n\n") {
		monkeys = append(monkeys, parseMonkey(monkeyStr))
	}
	return iterate(10000, monkeys, 1)
}

func Part2() {
	name := "Day #11 - part 2"

	common.TestOutputBig(name+" - input 2", 2713310158, Solve2(Input1))
	common.PrintOutputBig(name, Solve2(common.Readfile("./day11/input.txt")))
}
