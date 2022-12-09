package day09

import (
	"aoc/common"
	"math"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

func (c Coords) Up(amount int) Coords {
	return Coords{c.x, c.y + amount}
}
func (c Coords) Down(amount int) Coords {
	return Coords{c.x, c.y - amount}
}
func (c Coords) Right(amount int) Coords {
	return Coords{c.x + amount, c.y}
}
func (c Coords) Left(amount int) Coords {
	return Coords{c.x - amount, c.y}
}
func (c Coords) Previous(direction string) Coords {
	switch direction {
	case "U":
		return c.Down(1)
	case "D":
		return c.Up(1)
	case "R":
		return c.Left(1)
	case "L":
		return c.Right(1)
	}
	return c
}
func (c Coords) isTouching(c2 Coords) bool {
	xDistance := math.Abs(float64(c.x) - float64(c2.x))
	yDistance := math.Abs(float64(c.y) - float64(c2.y))

	return xDistance <= 1 && yDistance <= 1
}

func iterate(lines []string) (count int) {
	visited := make(map[Coords]bool)
	head := Coords{0, 0}
	tail := Coords{0, 0}
	count = 1
	visited[tail] = true
	for _, line := range lines {
		split := strings.Split(line, " ")
		amount, _ := strconv.Atoi(split[1])

		for a := 0; a < amount; a++ {
			switch split[0] {
			case "U":
				head = head.Up(1)
			case "D":
				head = head.Down(1)
			case "R":
				head = head.Right(1)
			case "L":
				head = head.Left(1)
			}
			if !tail.isTouching(head) {
				tail = head.Previous(split[0])
				if !visited[tail] {
					count++
				}
				visited[tail] = true
			}
		}
	}
	return count
}

func Solve1(input string) (count int) {
	return iterate(strings.Split(input, "\n"))
}

func Part1() {
	name := "Day #09 - part 1"

	common.TestOutput(name+" - input 1", 13, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day09/input.txt")))
}
