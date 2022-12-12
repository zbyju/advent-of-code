package day12

import (
	"aoc/common"
	"errors"
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

func getNext(lines []string, x Coords, canMove func(byte, byte) bool) []Coords {
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

func bfs(lines []string, start Coords, isEnd func(Coords) bool, canMove func(byte, byte) bool) (count int) {
	visited := make(map[Coords]int)
	q := []Coords{}
	var curr Coords

	q = append(q, start)
	visited[start] = 0

	for len(q) > 0 {
		curr = q[0]
		q = q[1:]
		v := visited[curr]

		next := getNext(lines, curr, canMove)

		for _, n := range next {
			if isEnd(n) {
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

func parseInput(input string, startChar rune, endChar rune) (lines []string, start Coords, end Coords) {
	lines = []string{}
	line := ""
	row := 0
	col := 0
	for _, c := range input {
		if c == '\n' {
			lines = append(lines, line)
			line = ""
			col = 0
			row++
			continue
		}
		if c == startChar {
			start = Coords{row, col}
		}
		if c == endChar {
			end = Coords{row, col}
		}
		if c == 'S' {
			c = 'a'
		} else if c == 'E' {
			c = 'z'
		}
		line += string(c)
		col++
	}
	lines = append(lines, line)
	return lines, start, end
}

func Solve1(input string) int {
	lines, start, end := parseInput(input, 'S', 'E')
	canMove := func(from, to byte) bool {
		return int(to)-int(from) <= 1
	}
	isEnd := func(b Coords) bool {
		return b.row == end.row && b.col == end.col
	}

	return bfs(lines, start, isEnd, canMove)
}

func Part1() {
	name := "Day #12 - part 1"

	common.TestOutput(name+" - input 1", 31, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day12/input.txt")))
}
