package day16

import (
	"aoc/common"
	"strings"
)

func Solve2(input string) (count int) {
	valves, flowRates, _ := parseInput(strings.Split(input, "\n"))
	graph := buildGraph(valves)
	trimmedGraph, trimmedValves := trimGraph(graph, flowRates)
	cache := make(map[Input]int)
	b := (1 << len(trimmedValves)) - 1
	max := 0
	for i := 0; i <= (b+1)/2; i++ {
		if m := run(&trimmedGraph, &trimmedValves, i, "AA", 26, &cache) + run(&trimmedGraph, &trimmedValves, b^i, "AA", 26, &cache); m > max {
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
