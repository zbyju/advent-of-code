package day20

import (
	"aoc/common"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func parseInput(lines []string) ([]int, []int) {
	nums := []int{}
	indices := []int{}
	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
		indices = append(indices, i)
	}
	return nums, indices
}

func modularIndex(i, mod int) int {
	if i < 0 {
		return i + mod
	}
	return i
}

func Solve1(input string) int {
	nums, indices := parseInput(strings.Split(input, "\n"))

	for i := range nums {
		j := slices.Index(indices, i)
		indices = slices.Delete(indices, j, j+1)
		indices = slices.Insert(indices, modularIndex((j+nums[i])%len(indices), len(indices)), i)
	}
	zero := slices.Index(indices, slices.Index(nums, 0))

	v1 := nums[indices[modularIndex((zero+1000)%len(nums), len(nums))]]
	v2 := nums[indices[modularIndex((zero+2000)%len(nums), len(nums))]]
	v3 := nums[indices[modularIndex((zero+3000)%len(nums), len(nums))]]

	return v1 + v2 + v3
}

func Part1() {
	name := "Day #20 - part 1"

	common.TestOutput(name+" - input 1", 3, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day20/input.txt")))
}
