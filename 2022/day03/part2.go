package day03

import (
	"aoc/common"
	"strings"
)

func makeGroups(lines []string) [][3]string {
	groups := [][3]string{}
	for i := 0; i < len(lines); i += 3 {
		group := [3]string{
			lines[i],
			lines[i+1],
			lines[i+2],
		}
		groups = append(groups, group)
	}
	return groups
}

func findBadge(group [3]string) rune {
	seen := make(map[rune]int)
	for i, elf := range group {
		for _, c := range elf {
			if seen[c] == i {
				seen[c] += 1
				if seen[c] == 3 {
					return c
				}
			}
		}
	}
	return '?'
}

func Solve2(input string) (sum int) {
	lines := strings.Split(input, "\n")
	groups := makeGroups(lines)

	for _, group := range groups {
		badge := findBadge(group)
		sum += LetterToScore(badge)
	}

	return sum
}

func Part2() {
	name := "Day #03 - part 1"

	common.TestOutput(name+" - input 1", 70, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day03/input.txt")))
}
