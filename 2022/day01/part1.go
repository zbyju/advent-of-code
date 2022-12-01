package main

import (
	"log"
	"os"
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
	max := 0
	for _, line := range lines {
		if len(line) > 0 {
			calories, _ := strconv.Atoi(line)
			elf += calories
			if elf > max {
				max = elf
			}
		} else {
			elf = 0
		}
	}

	return max
}

func main() {
	output1 := solve(input1)
	if output1 != 24000 {
		log.Fatalf("Input 1 should be 24000 but is: %d", output1)
	} else {
		log.Printf("Input 1 is ok: %d", output1)
	}

	buf, _ := os.ReadFile("input.txt")
	input := string(buf)
	output := solve(input)
	log.Printf("Part 1 output is: %d", output)
}
