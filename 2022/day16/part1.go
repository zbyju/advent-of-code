package day16

import (
	"aoc/common"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type RawValve struct {
	flowRate int
	next     []string
}

type Valve struct {
	index    int
	name     string
	flowRate int
	next     map[string]int
}

type Valve2 struct {
	flowRate int
	next     []int
}

func parseInput(lines []string) (_ map[string]RawValve, _ map[string]int, totalFlowRate int) {
	re := regexp.MustCompile(`Valve ([A-Z][A-Z]) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z][A-Z].*)`)
	valves := make(map[string]RawValve)
	flowRates := make(map[string]int)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		fr, _ := strconv.Atoi(matches[2])
		totalFlowRate += fr

		valve := RawValve{fr, strings.Split(matches[3], ", ")}
		flowRates[matches[1]] = fr
		valves[matches[1]] = valve
	}
	return valves, flowRates, totalFlowRate
}

func addPath(tunnels *map[string]Valve, from string, to string, distance int) {
	if from == to {
		return
	}
	tf := (*tunnels)[from]
	if tf.next == nil {
		tf.next = map[string]int{}
	}
	tf.next[to] = distance
}

func shortestPath(valves map[string]RawValve, tunnels *map[string]Valve, from string) {
	q := []string{from}
	visited := make(map[string]bool)
	curr := from
	for len(q) > 0 {
		curr = q[0]
		q = q[1:]
		visited[curr] = true

		for _, nextName := range valves[curr].next {
			addPath(tunnels, curr, nextName, 1)

			distance := (*tunnels)[from].next[curr] + 1
			if (*tunnels)[from].next[nextName] == 0 || (*tunnels)[from].next[nextName] > distance {
				addPath(tunnels, from, nextName, (*tunnels)[from].next[curr]+1)
			}
			if !visited[nextName] {
				q = append(q, nextName)
			}
		}
	}
}

func buildGraph(valves map[string]RawValve) map[string]Valve {
	tunnels := make(map[string]Valve)
	for from, valve := range valves {
		tunnels[from] = Valve{-1, from, valve.flowRate, make(map[string]int)}
		shortestPath(valves, &tunnels, from)
	}
	return tunnels
}

func trimGraph(graph map[string]Valve, flowRates map[string]int) []Valve2 {
	trimmedGraph := make(map[string]Valve)
	index := 0
	valveNames := []string{}
	for k := range flowRates {
		valveNames = append(valveNames, k)
	}
	sort.Slice(valveNames, func(i, j int) bool {
		return valveNames[i] < valveNames[j]
	})

	for _, valveName := range valveNames {
		fr := flowRates[valveName]
		if fr == 0 && valveName != "AA" {
			continue
		}
		trimmedGraph[valveName] = Valve{index, valveName, fr, make(map[string]int)}
		index++
		for to, distance := range graph[valveName].next {
			if flowRates[to] == 0 {
				continue
			}
			trimmedGraph[valveName].next[to] = distance
		}
	}
	valves := make([]Valve2, 16)
	for _, v := range trimmedGraph {
		next := make([]int, 16)
		for name, distance := range v.next {
			next[trimmedGraph[name].index] = distance
		}
		valves[v.index] = Valve2{v.flowRate, next}
	}
	return valves
}

type Input struct {
	bitmask     int
	curr        int
	minutesLeft int
}

func run(graph []Valve2, bitmask int, curr int, minutesLeft int, cache *map[Input]int) int {
	if max, ok := (*cache)[Input{bitmask, curr, minutesLeft}]; ok {
		return max
	}

	maxVal := 0
	for nextIndex, distance := range graph[curr].next {
		if distance == 0 {
			continue
		}
		time := minutesLeft - distance - 1
		if time <= 1 {
			continue
		}
		bit := 1 << nextIndex
		if bitmask&bit != 0 {
			continue
		}
		fr := graph[nextIndex].flowRate
		bm := bitmask | bit

		if m := run(graph, bm, nextIndex, time, cache) + fr*time; m > maxVal {
			maxVal = m
		}
	}
	(*cache)[Input{bitmask, curr, minutesLeft}] = maxVal
	return maxVal
}

func Solve1(input string) (count int) {
	rawValves, flowRates, _ := parseInput(strings.Split(input, "\n"))
	graph := buildGraph(rawValves)
	valves := trimGraph(graph, flowRates)
	cache := make(map[Input]int)
	max := run(valves, 0, 0, 30, &cache)
	return max
}

func Part1() {
	name := "Day #16 - part 1"

	common.TestOutput(name+" - input 1", 1651, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day16/input.txt")))
}
