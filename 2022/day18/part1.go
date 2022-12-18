package day18

import (
	"aoc/common"
	"strconv"
	"strings"
)

type Cube [3]float32

func Solve1(input string) (count int) {
	lines := strings.Split(input, "\n")
	sides := make(map[Cube]int)
	offsets := []Cube{{0, 0, 0.5}, {0, 0.5, 0}, {0.5, 0, 0}, {0, 0, -0.5}, {0, -0.5, 0}, {-0.5, 0, 0}}

	for _, line := range lines {
		split := strings.Split(line, ",")
		cube := Cube{}
		for i, s := range split {
			area, _ := strconv.Atoi(s)
			cube[i] = float32(area)
		}
		for _, offset := range offsets {
			key := Cube{cube[0] + offset[0], cube[1] + offset[1], cube[2] + offset[2]}
			if _, ok := sides[key]; !ok {
				sides[key] = 0
			}
			sides[key]++
		}
	}
	for _, v := range sides {
		if v == 1 {
			count++
		}
	}
	return count
}

func Part1() {
	name := "Day #18 - part 1"

	common.TestOutput(name+" - input 1", 64, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day18/input.txt")))
}
