package main

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

func solve(input string) int {
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

func main() {
	output1 := solve(input1)
	if output1 != 45000 {
		log.Fatalf("Input 1 should be 45000 but is: %d", output1)
	} else {
		log.Printf("Input 1 is ok: %d", output1)
	}

	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	output := solve(input)
	log.Printf("Part 1 output is: %d", output)
}
