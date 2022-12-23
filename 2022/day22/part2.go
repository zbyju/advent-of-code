package day22

import (
	"aoc/common"
)

func (c Coords) next2(grid []string, dir int) Coords {
	switch dir {
	case 0:
		return Coords{c.x + 1, c.y}
	case 1:
		return Coords{c.x, c.y + 1}
	case 2:
		return Coords{c.x - 1, c.y}
	case 3:
		return Coords{c.x, c.y - 1}
	}
	return Coords{}
}

func Solve2(input string) int {
	grid, cmds, pos := parseInput(input)
	for _, cmd := range cmds {
		for mv := 0; mv < cmd.move; mv++ {
			n := pos.toCoords().next2(grid, pos.dir)
			copyPos := pos.dir

			if n.y < 0 && 50 <= n.x && n.x < 100 && pos.dir == 3 {
				pos.dir = 0
				n.y, n.x = n.x+100, 0
			} else if n.x < 0 && 150 <= n.y && n.y < 200 && pos.dir == 2 {
				pos.dir = 1
				n.y, n.x = 0, n.y-100
			} else if n.y < 0 && 100 <= n.x && n.x < 150 && pos.dir == 3 {
				n.y, n.x = 199, n.x-100
			} else if n.y >= 200 && 0 <= n.x && n.x < 50 && pos.dir == 1 {
				n.y, n.x = 0, n.x+100
			} else if n.x >= 150 && 0 <= n.y && n.y < 50 && pos.dir == 0 {
				pos.dir = 2
				n.y, n.x = 149-n.y, 99
			} else if n.x == 100 && 100 <= n.y && n.y < 150 && pos.dir == 0 {
				pos.dir = 2
				n.y, n.x = 149-n.y, 149
			} else if n.y == 50 && 100 <= n.x && n.x < 150 && pos.dir == 1 {
				pos.dir = 2
				n.y, n.x = n.x-50, 99
			} else if n.x == 100 && 50 <= n.y && n.y < 100 && pos.dir == 0 {
				pos.dir = 3
				n.y, n.x = 49, n.y+50
			} else if n.y == 150 && 50 <= n.x && n.x < 100 && pos.dir == 1 {
				pos.dir = 2
				n.y, n.x = n.x+100, 49
			} else if n.x == 50 && 150 <= n.y && n.y < 200 && pos.dir == 0 {
				pos.dir = 3
				n.y, n.x = 149, n.y-100
			} else if n.y == 99 && 0 <= n.x && n.x < 50 && pos.dir == 3 {
				pos.dir = 0
				n.y, n.x = n.x+50, 50
			} else if n.x == 49 && 50 <= n.y && n.y < 100 && pos.dir == 2 {
				pos.dir = 1
				n.y, n.x = 100, n.y-50
			} else if n.x == 49 && 0 <= n.y && n.y < 50 && pos.dir == 2 {
				pos.dir = 0
				n.y, n.x = 149-n.y, 0
			} else if n.x < 0 && 100 <= n.y && n.y < 150 && pos.dir == 2 {
				pos.dir = 0
				n.y, n.x = 149-n.y, 50
			}

			if grid[n.y][n.x] == '#' {
				pos.dir = copyPos
				break
			}
			pos = Pos{n.x, n.y, pos.dir}
		}
		pos.dir = mod(pos.dir+cmd.dir, 4)
	}
	return (pos.y+1)*1000 + (pos.x+1)*4 + pos.dir
}

func Part2() {
	name := "Day #22 - part 2"

	common.PrintOutput(name, Solve2(common.Readfile("./day22/input.txt")))
}
