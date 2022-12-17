package day17

import (
	"aoc/common"
)

type Input struct {
	rocks    [PATTERN_SIZE][7]bool
	time     int
	numRocks int
}

type Cached struct {
	numRocks int
	height   int
}

const PATTERN_SIZE = 20

func runSimulation2(input string, maxRocks int) int {
	rocks := [][7]bool{}
	height := -1
	numRocks := 0
	time := 0
	cache := make(map[Input]Cached)
	offset := 0
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

		if cache != nil && len(rocks) > PATTERN_SIZE {
			key := Input{
				*(*[20][7]bool)(rocks[len(rocks)-20:]),
				time % len(input),
				numRocks % 5,
			}

			if cached, ok := cache[key]; ok {
				toGo := maxRocks - numRocks
				repetitions := toGo / (numRocks - cached.numRocks)
				offset = repetitions * (height - cached.height)
				numRocks += repetitions * (numRocks - cached.numRocks)
				cache = nil
			}
			if cache != nil {
				cache[key] = Cached{numRocks, height}
			}
		}
	}

	return height + offset
}

func Solve2(input string) int {
	maxHeight := runSimulation2(input, 1_000_000_000_000)
	return maxHeight + 1
}

func Part2() {
	name := "Day #17 - part 2"

	common.TestOutput(name+" - input 2", 1514285714288, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day17/input.txt")))
}
