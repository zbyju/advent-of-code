package day15

import (
	"aoc/common"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Block struct {
	sensor   Coords
	beacon   Coords
	distance int
}

type Coords struct {
	x int
	y int
}

type Interval struct {
	from int
	to   int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (c1 Coords) distance(c2 Coords) int {
	return abs(c1.x-c2.x) + abs(c1.y-c2.y)
}

func parseInput(lines []string, y int) ([]Block, map[int]bool) {
	re := regexp.MustCompile(`-?\d+`)
	blocks := []Block{}
	beaconsBlocking := make(map[int]bool)
	for _, line := range lines {
		parsed := re.FindAllStringSubmatch(line, -1)
		sx, _ := strconv.Atoi(parsed[0][0])
		sy, _ := strconv.Atoi(parsed[1][0])
		bx, _ := strconv.Atoi(parsed[2][0])
		by, _ := strconv.Atoi(parsed[3][0])
		sensor, beacon := Coords{sx, sy}, Coords{bx, by}
		blocks = append(blocks, Block{sensor, beacon, sensor.distance(beacon)})
		if beacon.y == y {
			beaconsBlocking[beacon.x] = true
		}
	}
	return blocks, beaconsBlocking
}

func Solve1(input string, y int) (count int) {
	intervals := []Interval{}
	blocks, beaconsBlocking := parseInput(strings.Split(input, "\n"), y)
	// Create intervals
	for _, block := range blocks {
		dx := block.distance - abs(block.sensor.y-y)
		if dx <= 0 {
			continue
		}

		lo := block.sensor.x - dx
		hi := block.sensor.x + dx
		intervals = append(intervals, Interval{lo, hi})
	}

	// Merge intervals
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].from == intervals[j].from {
			return intervals[i].to < intervals[j].to
		}
		return intervals[i].from < intervals[j].from
	})
	q := []Interval{}
	for _, interval := range intervals {
		if len(q) == 0 {
			q = append(q, interval)
			continue
		}

		if interval.from > q[len(q)-1].to+1 {
			q = append(q, interval)
			continue
		}
		q[len(q)-1].to = max(q[len(q)-1].to, interval.to)
	}

	// Count impossible positions
	beacons := 0
	for _, interval := range q {
		for b := range beaconsBlocking {
			if b >= interval.from && b <= interval.to {
				beacons++
			}
		}
		count += interval.to - interval.from + 1
	}
	return count - beacons
}

func Part1() {
	name := "Day #15 - part 1"

	common.TestOutput(name+" - input 1", 26, Solve1(Input1, 10))
	common.PrintOutput(name, Solve1(common.Readfile("./day15/input.txt"), 2_000_000))
}
