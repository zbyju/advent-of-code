package day18

import (
	"aoc/common"
	"math"
	"strconv"
	"strings"
)

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Solve2(input string) (count int) {
	lines := strings.Split(input, "\n")
	sides := make(map[Cube]int)
	offsets := []Cube{{0, 0, 0.5}, {0, 0.5, 0}, {0.5, 0, 0}, {0, 0, -0.5}, {0, -0.5, 0}, {-0.5, 0, 0}}
	mx, my, mz := float32(math.Inf(1)), float32(math.Inf(1)), float32(math.Inf(1))
	Mx, My, Mz := float32(math.Inf(-1)), float32(math.Inf(-1)), float32(math.Inf(-1))
	droplet := make(map[Cube]bool)

	for _, line := range lines {
		split := strings.Split(line, ",")
		cube := Cube{}
		for i, s := range split {
			area, _ := strconv.Atoi(s)
			cube[i] = float32(area)
		}
		droplet[cube] = true

		mx = min(mx, cube[0])
		my = min(my, cube[1])
		mz = min(mz, cube[2])

		Mx = max(Mx, cube[0])
		My = max(My, cube[1])
		Mz = max(Mz, cube[2])

		for _, offset := range offsets {
			key := Cube{cube[0] + offset[0], cube[1] + offset[1], cube[2] + offset[2]}
			if _, ok := sides[key]; !ok {
				sides[key] = 0
			}
			sides[key]++
		}
	}
	mx--
	my--
	mz--
	Mx++
	My++
	Mz++

	air := make(map[Cube]bool)
	air[Cube{mx, my, mz}] = true
	q := []Cube{{mx, my, mz}}

	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		for _, offset := range offsets {
			k := Cube{c[0] + offset[0]*2, c[1] + offset[1]*2, c[2] + offset[2]*2}

			if !(mx <= k[0] && k[0] <= Mx && my <= k[1] && k[1] <= My && mz <= k[2] && k[2] <= Mz) {
				continue
			}
			inDroplet := droplet[k]
			inAir := air[k]
			if inDroplet || inAir {
				continue
			}
			air[k] = true
			q = append(q, k)
		}
	}

	free := make(map[Cube]bool)
	for cube := range air {
		for _, offset := range offsets {
			free[Cube{cube[0] + offset[0], cube[1] + offset[1], cube[2] + offset[2]}] = true
		}
	}
	for k := range sides {
		if free[k] {
			count++
		}
	}
	return count
}

func Part2() {
	name := "Day #18 - part 2"

	common.TestOutput(name+" - input 2", 58, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day18/input.txt")))
}
