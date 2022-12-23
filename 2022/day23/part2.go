package day23

import (
	"aoc/common"
)

func Solve2(input string) int {
	elves := parseInput(input)

	round := 0
	for {
		// Propose moving
		propositions := make(map[Coords]int)
		stayed := 0
		for _, elf := range elves {
			proposing, didMove := elf.next(elves, round)
			if !didMove {
				stayed++
				if stayed == len(elves) {
					return round + 1
				}
			}
			elf.proposed = proposing
			propositions[proposing] += 1
		}

		// Check propositions
		newElves := make(map[Coords]*Elf)
		for _, elf := range elves {
			if propositions[elf.proposed] == 1 {
				elf.x = elf.proposed.x
				elf.y = elf.proposed.y
				elf.proposed = Coords{}
			}
			newElves[Coords{elf.x, elf.y}] = elf
		}

		elves = newElves
		round++
	}
}

func Part2() {
	name := "Day #23 - part 2"

	common.TestOutput(name+" - input 1", 20, Solve2(Input2))
	common.PrintOutput(name, Solve2(common.Readfile("./day23/input.txt")))
}
