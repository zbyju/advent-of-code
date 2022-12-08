package day08

import (
	"aoc/common"
)

func calcRowScore(trees [][]int, p Coords) int {
	left, right := 0, 0
	t := trees[p.row][p.col]

	// Left
	for i := p.col - 1; i >= 0; i-- {
		tree := trees[p.row][i]
		left++
		if tree >= t {
			break
		}
	}

	// Right
	for i := p.col + 1; i < len(trees[0]); i++ {
		tree := trees[p.row][i]
		right++
		if tree >= t {
			break
		}
	}

	return left * right
}

func calcColScore(trees [][]int, p Coords) int {
	up, down := 0, 0
	t := trees[p.row][p.col]

	// Up
	for i := p.row - 1; i >= 0; i-- {
		tree := trees[i][p.col]
		up++
		if tree >= t {
			break
		}
	}

	// Down
	for i := p.row + 1; i < len(trees); i++ {
		tree := trees[i][p.col]
		down++
		if tree >= t {
			break
		}
	}

	return down * up
}

func calcScore(trees [][]int, p Coords) int {
	if p.row == 0 || p.row == len(trees)-1 || p.col == 0 || p.col == len(trees[0])-1 {
		return 0
	}
	return calcRowScore(trees, p) * calcColScore(trees, p)
}

func Solve2(input string) (max int) {
	trees := parseInput(input)
	for i := range trees {
		for j := range trees[i] {
			if score := calcScore(trees, Coords{i, j}); score > max {
				max = score
			}
		}
	}

	return max
}

func Part2() {
	name := "Day #08 - part 2"

	common.TestOutput(name+" - input 1", 8, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day08/input.txt")))
}
