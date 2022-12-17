package day16

import (
	"aoc/common"
	"regexp"
	"strconv"
	"strings"
)

type Valve struct {
	flowRate int
	next     []string
}

type Tunnel struct {
	to map[string]int
}

type Val struct {
	index    int
	flowRate int
}

func parseInput(lines []string) (_ map[string]Valve, _ map[string]int, totalFlowRate int) {
	re := regexp.MustCompile(`Valve ([A-Z][A-Z]) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z][A-Z].*)`)
	valves := make(map[string]Valve)
	flowRates := make(map[string]int)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		fr, _ := strconv.Atoi(matches[2])
		totalFlowRate += fr

		valve := Valve{fr, strings.Split(matches[3], ", ")}
		flowRates[matches[1]] = fr
		valves[matches[1]] = valve
	}
	return valves, flowRates, totalFlowRate
}

func addPath(tunnels *map[string]Tunnel, from string, to string, distance int) {
	if from == to {
		return
	}
	tf := (*tunnels)[from]
	if tf.to == nil {
		tf.to = map[string]int{}
	}
	tf.to[to] = distance
}

func shortestPath(valves map[string]Valve, tunnels *map[string]Tunnel, from string) {
	q := []string{from}
	visited := make(map[string]bool)
	curr := from
	for len(q) > 0 {
		curr = q[0]
		q = q[1:]
		visited[curr] = true

		for _, nextName := range valves[curr].next {
			addPath(tunnels, curr, nextName, 1)

			distance := (*tunnels)[from].to[curr] + 1
			if (*tunnels)[from].to[nextName] == 0 || (*tunnels)[from].to[nextName] > distance {
				addPath(tunnels, from, nextName, (*tunnels)[from].to[curr]+1)
			}
			if !visited[nextName] {
				q = append(q, nextName)
			}
		}
	}
}

func buildGraph(valves map[string]Valve) map[string]Tunnel {
	tunnels := make(map[string]Tunnel)
	for from := range valves {
		tunnels[from] = Tunnel{make(map[string]int)}
		shortestPath(valves, &tunnels, from)
	}
	return tunnels
}

type Input struct {
	bitmask     int
	curr        int
	minutesLeft int
}

func run(graph *map[string]Tunnel, valves *map[string]Val, bitmask int, curr string, minutesLeft int, cache *map[Input]int) int {
	if max, ok := (*cache)[Input{bitmask, (*valves)[curr].index, minutesLeft}]; ok {
		return max
	}

	maxVal := 0
	for next, distance := range (*graph)[curr].to {
		time := minutesLeft - distance - 1
		if time <= 1 {
			continue
		}
		bit := 1 << (*valves)[next].index
		if bitmask&bit != 0 {
			continue
		}
		fr := (*valves)[next].flowRate
		bm := bitmask | bit
		if m := run(graph, valves, bm, next, time, cache) + fr*time; m > maxVal {
			maxVal = m
		}
	}
	(*cache)[Input{bitmask, (*valves)[curr].index, minutesLeft}] = maxVal
	return maxVal
}

func trimGraph(graph map[string]Tunnel, flowRates map[string]int) (map[string]Tunnel, map[string]Val) {
	trimmedGraph := make(map[string]Tunnel)
	trimmedValves := make(map[string]Val)
	index := 0
	for valve, fr := range flowRates {
		if fr == 0 && valve != "AA" {
			continue
		}
		trimmedValves[valve] = Val{index, fr}
		index++
		trimmedGraph[valve] = Tunnel{make(map[string]int)}
		for to, distance := range graph[valve].to {
			if flowRates[to] == 0 {
				continue
			}
			trimmedGraph[valve].to[to] = distance
		}
	}
	return trimmedGraph, trimmedValves
}

func Solve1(input string) (count int) {
	valves, flowRates, _ := parseInput(strings.Split(input, "\n"))
	graph := buildGraph(valves)
	trimmedGraph, trimmedValves := trimGraph(graph, flowRates)
	cache := make(map[Input]int)
	max := run(&trimmedGraph, &trimmedValves, 0, "AA", 30, &cache)
	return max
}

func Part1() {
	name := "Day #16 - part 1"

	common.TestOutput(name+" - input 1", 1651, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day16/input.txt")))
}
