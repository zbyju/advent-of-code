package day12

import (
	"aoc/common"
	"errors"
	"strings"
)

type Coords struct {
	row int
	col int
}

func findStart(lines []string) Coords {
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'S' {
				return Coords{i, j}
			}
		}
	}
	return Coords{}
}

func valAtCoords(lines []string, x Coords) (byte, error) {
	if x.row < 0 || x.row >= len(lines) || x.col < 0 || x.col >= len(lines[x.row]) {
		return '?', errors.New("out of bounds")
	}
	return lines[x.row][x.col], nil
}

func canMove(from, to byte) bool {
	if from == 'S' {
		from = 'a'
	}
	if from == 'z' && to == 'E' {
		return true
	} else if to == 'E' {
		return false
	}
	return int(to)-int(from) <= 1
}

func getNext(lines []string, x Coords) []Coords {
	next := []Coords{}
	y, _ := valAtCoords(lines, x)
	left := Coords{x.row - 1, x.col}
	right := Coords{x.row + 1, x.col}
	up := Coords{x.row, x.col - 1}
	down := Coords{x.row, x.col + 1}

	if v, err := valAtCoords(lines, left); err == nil && canMove(y, v) {
		next = append(next, left)
	}
	if v, err := valAtCoords(lines, right); err == nil && canMove(y, v) {
		next = append(next, right)
	}
	if v, err := valAtCoords(lines, up); err == nil && canMove(y, v) {
		next = append(next, up)
	}
	if v, err := valAtCoords(lines, down); err == nil && canMove(y, v) {
		next = append(next, down)
	}
	return next
}

func bfs(lines []string, start Coords) (count int) {
	visited := make(map[Coords]int)
	q := []Coords{}
	var curr Coords

	q = append(q, start)
	visited[start] = 0

	for len(q) > 0 {
		curr = q[0]
		q = q[1:]
		v := visited[curr]

		next := getNext(lines, curr)

		for _, n := range next {
			if lines[n.row][n.col] == 'E' {
				return v + 1
			}
			_, ok := visited[n]
			if !ok {
				visited[n] = v + 1
				q = append(q, n)
			}
		}
	}
	return -1
}

func Solve1(input string) int {
	lines := strings.Split(input, "\n")
	start := findStart(lines)
	return bfs(lines, start)
}

func Part1() {
	name := "Day #12 - part 1"

	common.TestOutput(name+" - input 1", 31, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day12/input.txt")))
}
