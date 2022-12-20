package day19

import (
	"aoc/common"
	"math"
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
	maxSpend      Resources
	oreRobot      Robot
	clayRobot     Robot
	obsidianRobot Robot
	geodeRobot    Robot
}

type Resources struct {
	ore      int
	clay     int
	obsidian int
	geode    int
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
}

func MaxSlice(s []int) (max int) {
	for _, x := range s {
		if x > max {
			max = x
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

	oreMax := MaxSlice([]int{oreOre, clayOre, obsidianOre, geodeOre})
	clayMax := obsidianClay
	obsidianMax := geodeObsidian

	return Blueprint{id, Resources{oreMax, clayMax, obsidianMax, 0}, ore, clay, obsidian, geode}
}

func parseInput(lines []string) []Factory {
	fs := []Factory{}
	for i, line := range lines {
		bp := parseBlueprint(line, i+1)
		fs = append(fs, Factory{bp, State{Resources{0, 0, 0, 0}, Robots{1, 0, 0, 0}}})
	}
	return fs
}

type Key struct {
	resources Resources
	robots    Robots
	time      int
}

func (f Factory) createRobot(r Robot, addRobot [4]int) Factory {
	return Factory{
		blueprint: f.blueprint,
		state: State{
			resources: Resources{
				ore:      f.state.resources.ore - r.ore,
				clay:     f.state.resources.clay - r.clay,
				obsidian: f.state.resources.obsidian - r.obsidian,
				geode:    f.state.resources.geode,
			},
			robots: Robots{
				ore:      f.state.robots.ore + addRobot[0],
				clay:     f.state.robots.clay + addRobot[1],
				obsidian: f.state.robots.obsidian + addRobot[2],
				geode:    f.state.robots.geode + addRobot[3],
			},
		},
	}
}

func (f Factory) wait(minutes int, maxNeeded [3]int) Factory {
	return Factory{
		blueprint: f.blueprint,
		state: State{
			resources: Resources{
				ore:      f.state.resources.ore + Min(f.state.robots.ore*minutes, maxNeeded[0]),
				clay:     f.state.resources.clay + Min(f.state.robots.clay*minutes, maxNeeded[1]),
				obsidian: f.state.resources.obsidian + Min(f.state.robots.obsidian*minutes, maxNeeded[2]),
				geode:    f.state.resources.geode + f.state.robots.geode*minutes,
			},
			robots: f.state.robots,
		},
	}
}

func (f Factory) maxNeeded(remainingTime int) [3]int {
	return [3]int{f.blueprint.maxSpend.ore * remainingTime, f.blueprint.maxSpend.clay * remainingTime, f.blueprint.maxSpend.obsidian * remainingTime}
}

func timeToGet(currentAmount, wantedAmount, production int) int {
	if currentAmount > wantedAmount {
		return 0
	}
	return int(math.Ceil(float64(wantedAmount-currentAmount) / float64(production)))
}

func run(f Factory, time int, cache *map[Key]int) int {
	if time == 0 {
		return f.state.resources.geode
	}
	key := Key{f.state.resources, f.state.robots, time}
	if val, ok := (*cache)[key]; ok {
		return val
	}

	max := f.state.resources.geode + f.state.robots.geode*time

	// Try to make geode robot
	if f.state.robots.obsidian > 0 {
		wait := Max(timeToGet(f.state.resources.ore, f.blueprint.geodeRobot.ore, f.state.robots.ore), timeToGet(f.state.resources.obsidian, f.blueprint.geodeRobot.obsidian, f.state.robots.obsidian))
		if remainingTime := time - wait - 1; remainingTime > 0 {
			m := Max(max, run(f.wait(wait+1, f.maxNeeded(remainingTime)).createRobot(f.blueprint.geodeRobot, [4]int{0, 0, 0, 1}), remainingTime, cache))
			if m > max {
				max = m
			}
		}
	}

	// Try to make ore robot
	if f.state.robots.ore < f.blueprint.maxSpend.ore {
		wait := timeToGet(f.state.resources.ore, f.blueprint.oreRobot.ore, f.state.robots.ore)
		if remainingTime := time - wait - 1; remainingTime > 0 {
			m := Max(max, run(f.wait(wait+1, f.maxNeeded(remainingTime)).createRobot(f.blueprint.oreRobot, [4]int{1, 0, 0, 0}), remainingTime, cache))
			if m > max {
				max = m
			}
		}
	}

	// Try to make clay robot
	if f.state.robots.clay < f.blueprint.maxSpend.clay {
		wait := timeToGet(f.state.resources.ore, f.blueprint.clayRobot.ore, f.state.robots.ore)
		if remainingTime := time - wait - 1; remainingTime > 0 {
			m := Max(max, run(f.wait(wait+1, f.maxNeeded(remainingTime)).createRobot(f.blueprint.clayRobot, [4]int{0, 1, 0, 0}), remainingTime, cache))
			if m > max {
				max = m
			}
		}
	}

	// Try to make obsidian robot
	if f.state.robots.obsidian < f.blueprint.maxSpend.obsidian && f.state.robots.clay > 0 {
		wait := Max(
			timeToGet(f.state.resources.ore, f.blueprint.obsidianRobot.ore, f.state.robots.ore),
			timeToGet(f.state.resources.clay, f.blueprint.obsidianRobot.clay, f.state.robots.clay),
		)
		if remainingTime := time - wait - 1; remainingTime > 0 {
			m := Max(max, run(f.wait(wait+1, f.maxNeeded(remainingTime)).createRobot(f.blueprint.obsidianRobot, [4]int{0, 0, 1, 0}), remainingTime, cache))
			if m > max {
				max = m
			}
		}
	}

	(*cache)[key] = max
	return max
}

func Solve1(input string) (total int) {
	fs := parseInput(strings.Split(input, "\n"))
	for _, f := range fs {
		m := run(f, 24, &map[Key]int{})
		total += f.blueprint.id * m
	}
	return total
}

func Part1() {
	name := "Day #19 - part 1"

	common.TestOutput(name+" - input 1", 33, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day19/input.txt")))
}
