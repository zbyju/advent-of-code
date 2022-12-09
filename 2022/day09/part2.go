package day09

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

func initRope(length int) []Coords {
	rope := []Coords{}
	for i := 0; i < length; i++ {
		rope = append(rope, Coords{0, 0})
	}
	return rope
}

func dir(from, to int) int {
	if to == from {
		return 0
	} else if to > from {
		return 1
	}
	return -1
}

func iterate2(lines []string) (count int) {
	visited := make(map[Coords]bool)
	rope := initRope(10)
	visited[Coords{0, 0}] = true
	count = 1
	for _, line := range lines {
		fmt.Println("Going through: ", line)
		split := strings.Split(line, " ")
		amount, _ := strconv.Atoi(split[1])

		for a := 0; a < amount; a++ {
			switch split[0] {
			case "U":
				rope[0] = rope[0].Up(1)
			case "D":
				rope[0] = rope[0].Down(1)
			case "R":
				rope[0] = rope[0].Right(1)
			case "L":
				rope[0] = rope[0].Left(1)
			}

			for knot := 1; knot < len(rope); knot++ {

				if !rope[knot].isTouching(rope[knot-1]) {
					rope[knot] = Coords{
						rope[knot].x + dir(rope[knot].x, rope[knot-1].x),
						rope[knot].y + dir(rope[knot].y, rope[knot-1].y),
					}
				}
			}
			if !visited[rope[len(rope)-1]] {
				count++
			}
			visited[rope[len(rope)-1]] = true
		}
	}
	return count
}

func Solve2(input string) (max int) {
	return iterate2(strings.Split(input, "\n"))
}

func Part2() {
	name := "Day #09 - part 2"

	common.TestOutput(name+" - input 1", 1, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day09/input.txt")))
}
