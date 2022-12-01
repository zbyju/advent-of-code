package day01

import (
	"aoc/common"
	"sort"
	"strconv"
	"strings"
)

const input1 = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func Solve2(input string) int {
	lines := strings.Split(input, "\n")

	elf := 0
	elfs := []int{}
	for _, line := range lines {
		if len(line) > 0 {
			calories, _ := strconv.Atoi(line)
			elf += calories
		} else {
			elfs = append(elfs, elf)
			elf = 0
		}
	}
	if elf != 0 {
		elfs = append(elfs, elf)
	}

	sort.Slice(elfs, func(a, b int) bool { return elfs[a] > elfs[b] })

	return elfs[0] + elfs[1] + elfs[2]
}

func Part2() {
	name := "Day #01 - part 2"

	common.TestOutput(name+" - input 1", 45000, Solve2(input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day01/input.txt")))
}
