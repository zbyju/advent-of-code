package day06

import (
	"aoc/common"
)

func initSlidingWindow(input string, length int) []rune {
	runes := []rune(input)
	window := make([]rune, length)
	for i := 0; i < length; i++ {
		window[i] = runes[i]
	}
	return window
}

func isUnique(window []rune) bool {
	for i := 0; i < len(window)-1; i++ {
		for j := i + 1; j < len(window); j++ {
			if window[i] == window[j] {
				return false
			}
		}
	}
	return true
}

func Solve1(input string, length int) int {
	window := initSlidingWindow(input, length)
	for i, v := range input[length:] {
		if isUnique(window) {
			return i + length
		}
		window[i%length] = v
	}
	return len(input)
}

func Part1() {
	name := "Day #06 - part 1"

	common.TestOutput(name+" - input 1", 7, Solve1(Input1, 4))
	common.TestOutput(name+" - input 2", 5, Solve1(Input2, 4))
	common.TestOutput(name+" - input 3", 6, Solve1(Input3, 4))
	common.TestOutput(name+" - input 4", 10, Solve1(Input4, 4))
	common.TestOutput(name+" - input 5", 11, Solve1(Input5, 4))
	common.PrintOutput(name, Solve1(common.Readfile("./day06/input.txt"), 4))
}
