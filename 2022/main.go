package main

import (
	"aoc/day00"
	"aoc/day01"
	"fmt"
	"os"
	"strings"
)

var solvers = map[string]func(){
	"00-0": day00.Hello,
	"01-0": day01.Part1,
	"01-1": day01.Part2,
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
	solvers[day+"-"+part]()
	fmt.Println()
}

func main() {

	if len(os.Args) == 1 {
		printHeading("Running all days and all parts")

		for i := range solvers {
			split := strings.Split(i, "-")
			run(split[0], split[1])
		}
	}
}
