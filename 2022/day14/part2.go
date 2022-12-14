package day14

import (
	"aoc/common"
	"strings"
)

func Solve2(input string) (sum int) {
	blocks, abyss := parseInput(strings.Split(input, "\n"))
	sandCount := 0
	for { // Spawn new
		sand := Coords{500, 0}
		if blocks[sand] {
			return sandCount
		}
		for { // Fall
			if sand.y > abyss {
				blocks[sand] = true
				break
			} else if ns := (Coords{sand.x, sand.y + 1}); !blocks[ns] {
				sand = ns
			} else if ns := (Coords{sand.x - 1, sand.y + 1}); !blocks[ns] {
				sand = ns
			} else if ns := (Coords{sand.x + 1, sand.y + 1}); !blocks[ns] {
				sand = ns
			} else {
				sandCount++
				blocks[sand] = true
				break
			}
		}
	}
}

func Part2() {
	name := "Day #14 - part 2"

	common.TestOutput(name+" - input 2", 93, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day14/input.txt")))
}
