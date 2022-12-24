package day24

import (
	"aoc/common"
	"strings"
)

type Pos struct {
	r int
	c int
}

type Blizzards struct {
	l map[Pos]bool
	r map[Pos]bool
	u map[Pos]bool
	d map[Pos]bool
}

func parseInput(input string) (Blizzards, Pos) {
	blizzards := Blizzards{make(map[Pos]bool), make(map[Pos]bool), make(map[Pos]bool), make(map[Pos]bool)}
	lines := strings.Split(input, "\n")
	for r, line := range lines[1 : len(lines)-1] {
		for c, item := range line[1 : len(line)-1] {
			if item == '<' {
				blizzards.l[Pos{r, c}] = true
			}
			if item == '>' {
				blizzards.r[Pos{r, c}] = true
			}
			if item == '^' {
				blizzards.u[Pos{r, c}] = true
			}
			if item == 'v' {
				blizzards.d[Pos{r, c}] = true
			}
		}
	}
	return blizzards, Pos{len(lines) - 2, len(lines[0]) - 2}
}

type State struct {
	time int
	pos  Pos
}

type BlizzardTest struct {
	dir int // 0 = left, 1 == right, 2 == up, 3 == down
	pos Pos
}

func gcd(first, second int) int {
	if first == 0 {
		return second
	}
	return gcd(second%first, first)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func checkCollision(blizzards Blizzards, check Pos, dir int) bool {
	switch dir {
	case 0:
		return blizzards.l[check]
	case 1:
		return blizzards.r[check]
	case 2:
		return blizzards.u[check]
	case 3:
		return blizzards.d[check]
	}
	return false
}

func bfs(blizzards Blizzards, size Pos) int {
	q := []State{{0, Pos{-1, 0}}}
	ds := []Pos{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {0, 0}}
	ts := []BlizzardTest{{0, Pos{0, -1}}, {1, Pos{0, 1}}, {2, Pos{-1, 0}}, {3, Pos{1, 0}}}
	target := Pos{size.r, size.c - 1}
	lcm := size.r * size.c / gcd(size.r, size.c)
	seen := make(map[State]bool)

	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		for _, d := range ds {
			np := Pos{s.pos.r + d.r, s.pos.c + d.c}

			time := s.time + 1

			if np == target {
				return time
			}

			if (np.r < 0 || np.c < 0 || np.r >= size.r || np.c >= size.c) && np != (Pos{-1, 0}) {
				continue
			}

			collision := false
			for _, t := range ts {
				check := Pos{mod(np.r-t.pos.r*time, size.r), mod(np.c-t.pos.c*time, size.c)}
				if checkCollision(blizzards, check, t.dir) {
					collision = true
					break
				}
			}
			if !collision {
				key := State{mod(time, lcm), np}

				if seen[key] {
					continue
				}

				seen[key] = true
				q = append(q, State{time, np})
			}
		}
	}

	return -1
}

func Solve1(input string) int {
	blizzards, size := parseInput(input)
	return bfs(blizzards, size)
}

func Part1() {
	name := "Day #24 - part 1"

	common.TestOutput(name+" - input 1", 18, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day24/input.txt")))
}
