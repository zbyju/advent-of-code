package day01

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve1(input string) int {
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

func Part1() {
	output1 := Solve1(input1)
	if output1 != 24000 {
		log.Fatalf("Input 1 should be 24000 but is: %d", output1)
	} else {
		log.Printf("Input 1 is ok: %d", output1)
	}

	buf, _ := os.ReadFile("./day01/input.txt")
	input := string(buf)
	output := Solve1(input)
	log.Printf("Part 1 output is: %d", output)
}
