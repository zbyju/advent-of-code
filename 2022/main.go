package main

import (
	"aoc/day00"
	"aoc/day01"
	"aoc/day02"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

func initSolvers() {
	solvers.Set("00-1", day00.Hello)
	solvers.Set("01-1", day01.Part1)
	solvers.Set("01-2", day01.Part2)
	solvers.Set("02-1", day02.Part1)
	solvers.Set("02-2", day02.Part2)
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

func run(day, part string) {
	printHeading2("Running day #" + day + " - part " + part)
	solver, ok := solvers.Get(day + "-" + part)

	if ok == false {
		log.Fatalf("The combination of %s - %s does not exist", day, part)
	}

	solver()
	fmt.Println()
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

		for _, v := range solvers.Keys() {
			split := strings.Split(v, "-")
			run(split[0], split[1])
		}
	}

	// Run the selected day
	if len(os.Args) == 2 {
		day := os.Args[1]
		if !isDayValid(day) {
			log.Fatalf("Day ('%s') is not valid", day)
		}
		paddedDay := dayToTwoDigit(day)

		if paddedDay == "00" {
			run("00", "1")
		} else {
			run(paddedDay, "1")
			run(paddedDay, "2")
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
		run(paddedDay, part)
	}
}
