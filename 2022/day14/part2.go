package day14

import (
	"aoc/common"
	"strings"
)

func Solve2(input string) (sum int) {
	blocks, abyss := parseInput(strings.Split(input, "\n"))
	sandCount := 0
	s := []Coords{{500, 0}}
	for len(s) > 0 { // Spawn new
		sand := s[len(s)-1]
		s = s[:len(s)-1]
		for { // Fall
			if sand.y > abyss {
				blocks[sand] = true
				break
			}
			if ns := (Coords{sand.x, sand.y + 1}); !blocks[ns] {
				s = append(s, sand)
				sand = ns
			} else if ns := (Coords{sand.x - 1, sand.y + 1}); !blocks[ns] {
				s = append(s, sand)
				sand = ns
			} else if ns := (Coords{sand.x + 1, sand.y + 1}); !blocks[ns] {
				s = append(s, sand)
				sand = ns
			} else {
				sandCount++
				blocks[sand] = true
				break
			}
		}
	}
	return sandCount
}

func Part2() {
	name := "Day #14 - part 2"

	common.TestOutput(name+" - input 2", 93, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day14/input.txt")))
}
