package day08

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
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

func isVisibleRow(trees [][]int, position Coords) bool {
	left, right := true, true
	t := trees[position.x][position.y]
	for j, tree := range trees[position.x] {
		fmt.Println("Checking row")
		if tree >= t {
			if j < position.y {
				left = false
			} else if j > position.y {
				right = false
			}
		}
	}
	fmt.Println("Left, right:", left, right)
	return left || right
}

func isVisibleCol(trees [][]int, position Coords) bool {
	up, down := true, true
	t := trees[position.x][position.y]
	for i, row := range trees {
		tree := row[i]
		if tree >= t {
			if i < position.x {
				down = false
			} else if i > position.x {
				up = false
			}
		}
	}
	fmt.Println("Down, up:", down, up)
	return down || up
}

func isVisible(trees [][]int, position Coords) bool {
	if position.x == 0 || position.x == len(trees) || position.y == 0 || position.y == len(trees[position.x]) {
		return true
	}
	return isVisibleRow(trees, position) || isVisibleCol(trees, position)
}

func Solve1(input string) (count int64) {
	trees := parseInput(input)

	for i := range trees {
		for j := range trees {
			fmt.Println(i, j, trees[i][j], isVisible(trees, Coords{i, j}))
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
