package main

import (
	"aoc/day00"
	"aoc/day01"
	"fmt"
	"os"
	"strings"

	"github.com/elliotchance/orderedmap/v2"
)

var solvers = orderedmap.NewOrderedMap[string, func()]()

func initSolvers() {
	solvers.Set("00-0", day00.Hello)
	solvers.Set("01-1", day01.Part1)
	solvers.Set("01-2", day01.Part2)
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
	solver, _ := solvers.Get(day + "-" + part)
	solver()
	fmt.Println()
}

func main() {
	initSolvers()

	if len(os.Args) == 1 {
		printHeading("Running all days and all parts")

		for _, v := range solvers.Keys() {
			split := strings.Split(v, "-")
			run(split[0], split[1])
		}
	}
}
