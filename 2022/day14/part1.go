package day14

import (
	"aoc/common"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

type Line struct {
	from Coords
	to   Coords
}

func addToSet(blocks *map[Coords]bool, line Line) {
	if line.from.x == line.to.x {
		from := line.from.y
		to := line.to.y
		if line.to.y < from {
			from = line.to.y
			to = line.from.y
		}
		for i := from; i <= to; i++ {
			(*blocks)[Coords{line.from.x, i}] = true
		}
	} else {
		from := line.from.x
		to := line.to.x
		if line.to.x < from {
			from = line.to.x
			to = line.from.x
		}
		for i := from; i <= to; i++ {
			(*blocks)[Coords{i, line.from.y}] = true
		}
	}
}

func parseInput(lines []string) (map[Coords]bool, int) {
	blocks := make(map[Coords]bool)
	last := Coords{}
	maxCol := -1
	for _, line := range lines {
		paths := strings.Split(line, " -> ")
		for i, path := range paths {
			split := strings.Split(path, ",")
			row, _ := strconv.Atoi(split[0])
			col, _ := strconv.Atoi(split[1])
			coords := Coords{row, col}
			if i != 0 {
				addToSet(&blocks, Line{last, coords})
			}
			last = coords
			if col > maxCol {
				maxCol = col
			}
		}
	}
	return blocks, maxCol + 1
}

func Solve1(input string) (sum int) {
	blocks, abyss := parseInput(strings.Split(input, "\n"))
	sandCount := 0
	s := []Coords{{500, 0}}
	for { // Spawn new
		sand := s[len(s)-1]
		s = s[:len(s)-1]
		for { // Fall
			if ns := (Coords{sand.x, sand.y + 1}); !blocks[ns] {
				if ns.y >= abyss {
					return sandCount
				}
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
}

func Part1() {
	name := "Day #14 - part 1"

	common.TestOutput(name+" - input 1", 24, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day14/input.txt")))
}
