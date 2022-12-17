package day17

import (
	"aoc/common"
	"fmt"
)

type Coords struct {
	x int
	y int
}

func getJet(input string, time int) byte {
	return input[time%len(input)]
}

func spawnRock(rockNum, height int) []Coords {
	switch rockNum % 5 {
	case 0:
		return []Coords{{2, height + 3}, {3, height + 3}, {4, height + 3}, {5, height + 3}}
	case 1:
		return []Coords{{3, height + 3}, {2, height + 4}, {3, height + 4}, {4, height + 4}, {3, height + 5}}
	case 2:
		return []Coords{{2, height + 3}, {3, height + 3}, {4, height + 3}, {4, height + 4}, {4, height + 5}}
	case 3:
		return []Coords{{2, height + 3}, {2, height + 4}, {2, height + 5}, {2, height + 6}}
	case 4:
		return []Coords{{2, height + 3}, {3, height + 3}, {2, height + 4}, {3, height + 4}}
	}
	return []Coords{}
}

func isEmpty(rocks [][7]bool, c Coords) bool {
	if c.x >= 7 || c.x < 0 || c.y < 0 {
		return false
	}
	if c.y >= len(rocks) {
		return true
	}
	return !rocks[c.y][c.x]
}

func canFall(rocks [][7]bool, rock []Coords) bool {
	if rock[0].y < 0 {
		return false
	}
	for _, r := range rock {
		if !isEmpty(rocks, Coords{r.x, r.y - 1}) {
			return false
		}
	}
	return true
}

func pushRock(input string, i int, rocks [][7]bool, rock []Coords) []Coords {
	xChange := 1
	if getJet(input, i) == '<' {
		xChange = -1
	}
	newRock := []Coords{}
	for _, r := range rock {
		newCoords := Coords{r.x + xChange, r.y}
		if !isEmpty(rocks, newCoords) {
			return rock
		} else {
			newRock = append(newRock, newCoords)
		}
	}
	return newRock
}

func fall(rock []Coords) []Coords {
	newRock := []Coords{}
	for _, r := range rock {
		newRock = append(newRock, Coords{r.x, r.y - 1})
	}
	return newRock
}

func stopRock(rocks *[][7]bool, rock []Coords, height int) int {
	for _, r := range rock {
		if r.y > height {
			height = r.y
		}
		for r.y >= len(*rocks) {
			(*rocks) = append((*rocks), [7]bool{false, false, false, false, false, false, false})
		}
		(*rocks)[r.y][r.x] = true
	}
	return height
}

func print(rocks [][7]bool, rock []Coords) {
	stopRock(&rocks, rock, 0)
	for ry := range rocks {
		y := len(rocks) - 1 - ry
		fmt.Print("|")
		for _, r := range rocks[y] {
			if r {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("|")
	}
}

func runSimulation(input string, maxRocks int) int {
	rocks := [][7]bool{}
	height := -1
	numRocks := 0
	time := 0
	for numRocks < maxRocks {
		rock := spawnRock(numRocks, height+1)
		numRocks++

		rock = pushRock(input, time, rocks, rock)
		time++

		for canFall(rocks, rock) {
			rock = fall(rock)
			rock = pushRock(input, time, rocks, rock)
			time++
		}

		height = stopRock(&rocks, rock, height)
	}
	return height
}

func Solve1(input string) int {
	maxHeight := runSimulation(input, 2022)
	return maxHeight + 1
}

func Part1() {
	name := "Day #17 - part 1"

	common.TestOutput(name+" - input 1", 3068, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day17/input.txt")))
}
