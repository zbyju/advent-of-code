package day15

import (
	"aoc/common"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Solve2(input string, max int) (count int64) {
	blocks, _ := parseInput(strings.Split(input, "\n"), -1)
	negLines, posLines := []int{}, []int{}

	for _, block := range blocks {
		negLines = append(negLines, block.sensor.x+block.sensor.y-block.distance)
		negLines = append(negLines, block.sensor.x+block.sensor.y+block.distance)
		posLines = append(posLines, block.sensor.x-block.sensor.y-block.distance)
		posLines = append(posLines, block.sensor.x-block.sensor.y+block.distance)
	}

	var neg, pos int
	for i := 0; i < len(negLines); i++ {
		for j := i; j < len(negLines); j++ {
			if abs(negLines[i]-negLines[j]) == 2 {
				neg = min(negLines[i], negLines[j]) + 1
			}
			if abs(posLines[i]-posLines[j]) == 2 {
				pos = min(posLines[i], posLines[j]) + 1
			}
		}
	}
	x := (pos + neg) / 2
	y := (neg - pos) / 2
	return int64(x)*4_000_000 + int64(y)
}

func Part2() {
	name := "Day #15 - part 2"

	// common.TestOutputBig(name+" - input 2", 56000011, Solve2(Input1, 20))
	common.PrintOutputBig(name, Solve2(common.Readfile("./day15/input.txt"), 4_000_000))
}
