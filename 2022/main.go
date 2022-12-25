package main

import (
	"aoc/day00"
	"aoc/day01"
	"aoc/day02"
	"aoc/day03"
	"aoc/day04"
	"aoc/day05"
	"aoc/day06"
	"aoc/day07"
	"aoc/day08"
	"aoc/day09"
	"aoc/day10"
	"aoc/day11"
	"aoc/day12"
	"aoc/day13"
	"aoc/day14"
	"aoc/day15"
	"aoc/day16"
	"aoc/day17"
	"aoc/day18"
	"aoc/day19"
	"aoc/day20"
	"aoc/day21"
	"aoc/day22"
	"aoc/day23"
	"aoc/day24"
	"aoc/day25"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elliotchance/orderedmap/v2"
)

var solvers = orderedmap.NewOrderedMap[string, func()]()

func isDayValid(day string) bool {
	num, err := strconv.Atoi(day)

	if err != nil || num < 0 || num > 25 {
		return false
	}

	return len(day) == 1 || len(day) == 2
}

func isPartValid(part string) bool {
	return part == "1" || part == "2"
}

func dayToTwoDigit(day string) string {
	if len(day) == 1 {
		return "0" + day
	} else {
		return day
	}
}

func measureTime(solver func()) int64 {
	start := time.Now()

	solver()

	elapsed := time.Since(start)

	fmt.Printf("Elapsed time: %s", elapsed)

	return elapsed.Nanoseconds()
}

func initSolvers() {
	solvers.Set("00-1", day00.Hello)
	solvers.Set("01-1", day01.Part1)
	solvers.Set("01-2", day01.Part2)
	solvers.Set("02-1", day02.Part1)
	solvers.Set("02-2", day02.Part2)
	solvers.Set("03-1", day03.Part1)
	solvers.Set("03-2", day03.Part2)
	solvers.Set("04-1", day04.Part1)
	solvers.Set("04-2", day04.Part2)
	solvers.Set("05-1", day05.Part1)
	solvers.Set("05-2", day05.Part2)
	solvers.Set("06-1", day06.Part1)
	solvers.Set("06-2", day06.Part2)
	solvers.Set("07-1", day07.Part1)
	solvers.Set("07-2", day07.Part2)
	solvers.Set("08-1", day08.Part1)
	solvers.Set("08-2", day08.Part2)
	solvers.Set("09-1", day09.Part1)
	solvers.Set("09-2", day09.Part2)
	solvers.Set("10-1", day10.Part1)
	solvers.Set("10-2", day10.Part2)
	solvers.Set("11-1", day11.Part1)
	solvers.Set("11-2", day11.Part2)
	solvers.Set("12-1", day12.Part1)
	solvers.Set("12-2", day12.Part2)
	solvers.Set("13-1", day13.Part1)
	solvers.Set("13-2", day13.Part2)
	solvers.Set("14-1", day14.Part1)
	solvers.Set("14-2", day14.Part2)
	solvers.Set("15-1", day15.Part1)
	solvers.Set("15-2", day15.Part2)
	solvers.Set("16-1", day16.Part1)
	solvers.Set("16-2", day16.Part2)
	solvers.Set("17-1", day17.Part1)
	solvers.Set("17-2", day17.Part2)
	solvers.Set("18-1", day18.Part1)
	solvers.Set("18-2", day18.Part2)
	solvers.Set("19-1", day19.Part1)
	solvers.Set("19-2", day19.Part2)
	solvers.Set("20-1", day20.Part1)
	solvers.Set("20-2", day20.Part2)
	solvers.Set("21-1", day21.Part1)
	solvers.Set("21-2", day21.Part2)
	solvers.Set("22-1", day22.Part1)
	solvers.Set("22-2", day22.Part2)
	solvers.Set("23-1", day23.Part1)
	solvers.Set("23-2", day23.Part2)
	solvers.Set("24-1", day24.Part1)
	solvers.Set("24-2", day24.Part2)
	solvers.Set("25-1", day25.Part1)
	solvers.Set("25-2", day25.Part2)
}

func printHeading(text string) {
	fmt.Println(strings.Repeat("=", len(text)))
	fmt.Println(text)
	fmt.Println(strings.Repeat("=", len(text)))
	fmt.Println()
}

func printHeading2(text string) {
	fmt.Println(strings.Repeat("-", len(text)))
	fmt.Println(text)
	fmt.Println(strings.Repeat("-", len(text)))
}

func printTime(t int64) {
	fmt.Println()
	fmt.Printf("Total time to run: %s\n", time.Duration(t))
	fmt.Println()
}

func run(day, part string) int64 {
	printHeading2("Running day #" + day + " - part " + part)
	solver, ok := solvers.Get(day + "-" + part)

	if !ok {
		log.Fatalf("The combination of %s - %s does not exist", day, part)
	}

	t := measureTime(solver)
	fmt.Println()
	return t
}

/*
Program can be ran 3 different ways:
 1. go run .						- runs all the days
 2. go run . day				- runs the selected day
 3. go run . day part		- runs the selected part from the selected day
*/
func main() {
	initSolvers()

	// Run all
	if len(os.Args) == 1 {
		printHeading("Running all days and all parts")

		var totalTime int64 = 0
		for _, v := range solvers.Keys() {
			split := strings.Split(v, "-")
			t := run(split[0], split[1])
			totalTime += t
		}

		printTime(totalTime)
	}

	// Run the selected day
	if len(os.Args) == 2 {
		day := os.Args[1]
		if !isDayValid(day) {
			log.Fatalf("Day ('%s') is not valid", day)
		}
		paddedDay := dayToTwoDigit(day)

		if paddedDay == "00" {
			t := run("00", "1")
			printTime(t)
		} else {
			t1 := run(paddedDay, "1")
			t2 := run(paddedDay, "2")
			printTime(t1 + t2)
		}
	}

	// Run the selected day and selected part
	if len(os.Args) == 3 {
		day := os.Args[1]
		part := os.Args[2]
		if !isDayValid(day) {
			log.Fatalf("Day ('%s') is not valid", day)
		}
		if !isPartValid(part) {
			log.Fatalf("Part ('%s') is not valid", part)
		}

		paddedDay := dayToTwoDigit(day)
		t := run(paddedDay, part)
		printTime(t)
	}
}
