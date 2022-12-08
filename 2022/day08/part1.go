package day08

import (
	"aoc/common"
	"strconv"
	"strings"
)

type Coords struct {
	row int
	col int
}

func parseInput(input string) [][]int {
	trees := [][]int{}
	for _, line := range strings.Split(input, "\n") {
		row := []int{}
		for _, treeRune := range line {
			tree, _ := strconv.Atoi(string(treeRune))
			row = append(row, tree)
		}
		trees = append(trees, row)
	}
	return trees
}

func isVisibleRow(trees [][]int, p Coords) bool {
	left, right := true, true
	t := trees[p.row][p.col]
	for j, tree := range trees[p.row] {
		if tree >= t {
			if j < p.col {
				left = false
			} else if j > p.col {
				right = false
			}
		}
	}
	return left || right
}

func isVisibleCol(trees [][]int, p Coords) bool {
	up, down := true, true
	t := trees[p.row][p.col]
	for i, row := range trees {
		tree := row[p.col]
		if tree >= t {
			if i < p.row {
				up = false
			} else if i > p.row {
				down = false
			}
		}
	}
	return down || up
}

func isVisible(trees [][]int, p Coords) bool {
	if p.row == 0 || p.row == len(trees)-1 || p.col == 0 || p.col == len(trees[0])-1 {
		return true
	}
	return isVisibleRow(trees, p) || isVisibleCol(trees, p)
}

func Solve1(input string) (count int64) {
	trees := parseInput(input)
	for i := range trees {
		for j := range trees[i] {
			if isVisible(trees, Coords{i, j}) {
				count++
			}
		}
	}

	return count
}

func Part1() {
	name := "Day #08 - part 1"

	common.TestOutputBig(name+" - input 1", 21, Solve1(Input1))
	common.PrintOutputBig(name, Solve1(common.Readfile("./day08/input.txt")))
}
