package day04

import (
	"aoc/common"
	"strconv"
	"strings"
)

type Interval struct {
	start int
	end   int
}

func doesContain(interval1, interval2 Interval) bool {
	return interval1.start <= interval2.start && interval1.end >= interval2.end
}

func ParseInterval(str string) Interval {
	split := strings.Split(str, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])
	return Interval{start, end}
}

func ParseLine(line string) [2]Interval {
	split := strings.Split(line, ",")
	interval1 := ParseInterval(split[0])
	interval2 := ParseInterval(split[1])
	return [2]Interval{interval1, interval2}
}

func Solve1(input string) (count int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		sections := ParseLine(line)
		if doesContain(sections[0], sections[1]) || doesContain(sections[1], sections[0]) {
			count++
		}
	}
	return count
}

func Part1() {
	name := "Day #04 - part 1"

	common.TestOutput(name+" - input 1", 2, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day04/input.txt")))
}
