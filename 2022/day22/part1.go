package day22

import (
	"aoc/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Cmd struct {
	move int
	dir  int
}

type Coords struct {
	x int
	y int
}

type Pos struct {
	x   int
	y   int
	dir int
}

type Box struct {
	left  int
	right int
	top   int
	bot   int
}

func (p Pos) toCoords() Coords {
	return Coords{p.x, p.y}
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func parseInput(input string) ([]string, []Cmd, Pos) {
	split := strings.Split(input, "\n\n")
	grid := strings.Split(split[0], "\n")

	re := regexp.MustCompile(`(\d+)([RL]?)`)
	matches := re.FindAllStringSubmatch(split[1], -1)

	cmds := []Cmd{}
	for _, match := range matches {
		move, _ := strconv.Atoi(match[1])
		dir := 0
		if match[2] != "" {
			if match[2] == "R" {
				dir = 1
			} else {
				dir = -1
			}
		}
		cmds = append(cmds, Cmd{move, dir})
	}

	start := Pos{}
	for i, x := range grid[0] {
		if x == '.' {
			start = Pos{i, 0, 0}
			break
		}
	}

	return grid, cmds, start
}

func gridAt(grid []string, c Coords) rune {
	if c.x >= 0 && c.y >= 0 && c.y < len(grid) && c.x < len(grid[c.y]) {
		return rune(grid[c.y][c.x])
	}
	return ' '
}

func (c Coords) next(grid []string, dir int) Coords {
	switch dir {
	case 0:
		return Coords{mod(c.x+1, len(grid[c.y])), c.y}
	case 1:
		return Coords{c.x, mod(c.y+1, len(grid))}
	case 2:
		return Coords{mod(c.x-1, len(grid[c.y])), c.y}
	case 3:
		return Coords{c.x, mod(c.y-1, len(grid))}
	}
	return Coords{}
}

func Solve1(input string, size int) int {
	grid, cmds, pos := parseInput(input)
	for _, cmd := range cmds {
		for mv := 0; mv < cmd.move; mv++ {
			n := pos.toCoords().next(grid, pos.dir)
			// fmt.Println("S:", pos, n)
			for gridAt(grid, n) == ' ' {
				n = n.next(grid, pos.dir)
			}

			if grid[n.y][n.x] == '#' {
				break
			}
			fmt.Println("E:", pos, n)
			pos = Pos{n.x, n.y, pos.dir}
		}
		pos.dir = mod(pos.dir+cmd.dir, 4)
	}
	return (pos.y+1)*1000 + (pos.x+1)*4 + pos.dir
}

func Part1() {
	name := "Day #22 - part 1"

	common.TestOutput(name+" - input 1", 6032, Solve1(Input1, 4))
	common.PrintOutput(name, Solve1(common.Readfile("./day22/input.txt"), 50))
}
