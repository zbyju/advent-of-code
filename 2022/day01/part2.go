package day01

import (
	"log"
	"os"
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
	buf, _ := os.ReadFile("./day01/input.txt")
	input := string(buf)
	output := Solve2(input)
	log.Printf("Part 2 output is: %d", output)
}
