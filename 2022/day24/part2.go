package day24

import (
	"aoc/common"
)

type State2 struct {
	time  int
	pos   Pos
	stage int
}

func bfs2(blizzards Blizzards, size Pos) int {
	q := []State2{{0, Pos{-1, 0}, 0}}
	ds := []Pos{{0, 1}, {0, -1}, {-1, 0}, {1, 0}, {0, 0}}
	ts := []BlizzardTest{{0, Pos{0, -1}}, {1, Pos{0, 1}}, {2, Pos{-1, 0}}, {3, Pos{1, 0}}}

	targets := [2]Pos{{size.r, size.c - 1}, {-1, 0}}

	lcm := size.r * size.c / gcd(size.r, size.c)
	seen := make(map[State2]bool)

	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		for _, d := range ds {
			np := Pos{s.pos.r + d.r, s.pos.c + d.c}

			time := s.time + 1

			nstage := s.stage
			if np == targets[s.stage%2] {
				if s.stage == 2 {
					return time
				}
				nstage += 1
			}

			if (np.r < 0 || np.c < 0 || np.r >= size.r || np.c >= size.c) && np != targets[0] && np != targets[1] {
				continue
			}

			collision := false
			if np != targets[0] && np != targets[1] {
				for _, t := range ts {
					check := Pos{mod(np.r-t.pos.r*time, size.r), mod(np.c-t.pos.c*time, size.c)}
					if checkCollision(blizzards, check, t.dir) {
						collision = true
						break
					}
				}
			}
			if !collision {
				key := State2{mod(time, lcm), np, nstage}

				if seen[key] {
					continue
				}

				seen[key] = true
				q = append(q, State2{time, np, nstage})
			}
		}
	}

	return -1
}

func Solve2(input string) int {
	blizzards, size := parseInput(input)
	return bfs2(blizzards, size)
}

func Part2() {
	name := "Day #24 - part 2"

	common.TestOutput(name+" - input 1", 54, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day24/input.txt")))
}
