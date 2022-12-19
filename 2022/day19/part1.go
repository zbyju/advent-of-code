package day19

import (
	"aoc/common"
	"regexp"
	"strconv"
	"strings"
)

type Factory struct {
	blueprint Blueprint
	state     State
}

type Robot struct {
	ore      int
	clay     int
	obsidian int
}

type Blueprint struct {
	id            int
	oreRobot      Robot
	clayRobot     Robot
	obsidianRobot Robot
	geodeRobot    Robot
}

type Resources struct {
	ore      int
	clay     int
	obsidian int
}

type Robots struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

type State struct {
	resources Resources
	robots    Robots
	geode     int
}

func parseBlueprint(line string, id int) Blueprint {
	re := regexp.MustCompile(`(\d+) (\w+)`)
	matches := re.FindAllStringSubmatch(line, -1)
	oreOre, _ := strconv.Atoi(matches[0][1])
	clayOre, _ := strconv.Atoi(matches[1][1])
	obsidianOre, _ := strconv.Atoi(matches[2][1])
	obsidianClay, _ := strconv.Atoi(matches[3][1])
	geodeOre, _ := strconv.Atoi(matches[4][1])
	geodeObsidian, _ := strconv.Atoi(matches[5][1])

	ore := Robot{oreOre, 0, 0}
	clay := Robot{clayOre, 0, 0}
	obsidian := Robot{obsidianOre, obsidianClay, 0}
	geode := Robot{geodeOre, 0, geodeObsidian}

	return Blueprint{id, ore, clay, obsidian, geode}
}

func parseInput(lines []string) []Factory {
	fs := []Factory{}
	for i, line := range lines {
		bp := parseBlueprint(line, i+1)
		fs = append(fs, Factory{bp, State{Resources{0, 0, 0}, Robots{1, 0, 0, 0}, 0}})
	}
	return fs
}

func (f *Factory) createRobot(r Robot) {
	(*f).state.resources.ore -= r.ore
	(*f).state.resources.clay -= r.clay
	(*f).state.resources.obsidian -= r.obsidian
}

func (f *Factory) nextRound(r Robot) {

}

func run(f Factory) int {

}

func Solve1(input string) int {
	fs := parseInput(strings.Split(input, "\n"))
	max := 0
	for _, f := range fs {
		m := run(f)
		if m > max {
			max = m
		}
	}
	return max
}

func Part1() {
	name := "Day #19 - part 1"

	common.TestOutput(name+" - input 1", 33, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day19/input.txt")))
}
