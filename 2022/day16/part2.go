package day16

import (
	"aoc/common"
	"strings"
)

func Solve2(input string) (count int) {
	rawValves, flowRates, _ := parseInput(strings.Split(input, "\n"))
	graph := buildGraph(rawValves)
	valves := trimGraph(graph, flowRates)
	cache := make(map[Input]int)
	b := (1 << len(valves)) - 1
	max := 0
	for i := 1; i <= (b+1)/2; i += 2 {
		if m := run(valves, i, 0, 26, &cache) + run(valves, b^i, 0, 26, &cache); m > max {
			max = m
		}
	}
	return max
}
func Part2() {
	name := "Day #16 - part 2"

	common.TestOutput(name+" - input 2", 1707, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day16/input.txt")))
}
