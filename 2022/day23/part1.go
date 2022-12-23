package day23

import (
	"aoc/common"
	"strings"
)

type Coords struct {
	x int
	y int
}

type Elf struct {
	x        int
	y        int
	proposed Coords
}

func (elf Elf) toCoords() Coords {
	return Coords{elf.x, elf.y}
}

func (c Coords) move(vec Coords) Coords {
	return Coords{c.x + vec.x, c.y + vec.y}
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func parseInput(input string) map[Coords]*Elf {
	elves := make(map[Coords]*Elf)
	for y, row := range strings.Split(input, "\n") {
		for x, r := range row {
			if r == '#' {
				elves[Coords{x, y}] = &Elf{x, y, Coords{}}
			}
		}
	}
	return elves
}

func (elf Elf) adjacent() map[string]Coords {
	adjacent := make(map[string]Coords)
	vecs := make(map[string]Coords)

	vecs["N"] = Coords{0, -1}
	vecs["S"] = Coords{0, 1}
	vecs["W"] = Coords{-1, 0}
	vecs["E"] = Coords{1, 0}

	vecs["NE"] = Coords{1, -1}
	vecs["NW"] = Coords{-1, -1}
	vecs["SE"] = Coords{1, 1}
	vecs["SW"] = Coords{-1, 1}

	for dir, vec := range vecs {
		adjacent[dir] = elf.toCoords().move(vec)
	}
	return adjacent
}

func (elf Elf) adjacentOccupied(elves map[Coords]*Elf) map[string]bool {
	adjacent := elf.adjacent()
	occupied := make(map[string]bool)

	for name, c := range adjacent {
		occupied[name] = elves[c] != nil
	}
	return occupied
}

func (elf Elf) next(elves map[Coords]*Elf, round int) (Coords, bool) {
	occupied := elf.adjacentOccupied(elves)
	isAlone := true

	for _, o := range occupied {
		if o {
			isAlone = false
		}
	}

	if isAlone {
		return elf.toCoords(), false
	}

	// dir: 0 = N, 1 = S, 2 = W, 3 = E
	for d := 0; d < 4; d++ {
		dir := mod(round+d, 4)

		if dir == 0 && !occupied["N"] && !occupied["NE"] && !occupied["NW"] {
			return Coords{elf.x, elf.y - 1}, true
		}
		if dir == 1 && !occupied["S"] && !occupied["SE"] && !occupied["SW"] {
			return Coords{elf.x, elf.y + 1}, true
		}
		if dir == 2 && !occupied["W"] && !occupied["NW"] && !occupied["SW"] {
			return Coords{elf.x - 1, elf.y}, true
		}
		if dir == 3 && !occupied["E"] && !occupied["NE"] && !occupied["SE"] {
			return Coords{elf.x + 1, elf.y}, true
		}
	}
	return elf.toCoords(), false
}

func Solve1(input string) int {
	elves := parseInput(input)

	const rounds int = 10

	for round := 0; round < rounds; round++ {
		// Propose moving
		propositions := make(map[Coords]int)
		for _, elf := range elves {
			proposing, _ := elf.next(elves, round)
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
	}

	var minX, minY, maxX, maxY int
	first := true
	for _, elf := range elves {
		if first {
			minX = elf.x
			maxX = elf.x
			minY = elf.y
			maxY = elf.y
			first = false
		}
		if elf.x < minX {
			minX = elf.x
		}
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.y < minY {
			minY = elf.y
		}
		if elf.y > maxY {
			maxY = elf.y
		}
	}

	return (maxX-minX+1)*(maxY-minY+1) - len(elves)
}

func Part1() {
	name := "Day #23 - part 1"

	common.TestOutput(name+" - input 1", 25, Solve1(Input1))
	common.TestOutput(name+" - input 2", 110, Solve1(Input2))
	common.PrintOutput(name, Solve1(common.Readfile("./day23/input.txt")))
}
