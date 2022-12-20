package day20

import (
	"aoc/common"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func parseInput2(lines []string, mult int64) ([]int64, []int) {
	nums := []int64{}
	indices := []int{}
	for i, line := range lines {
		num, _ := strconv.ParseInt(line, 10, 64)
		nums = append(nums, num*mult)
		indices = append(indices, i)
	}
	return nums, indices
}

func modularIndex2(i, mod int64) int64 {
	if i < 0 {
		return i + mod
	}
	return i
}

func Solve2(input string) int64 {
	nums, indices := parseInput2(strings.Split(input, "\n"), 811589153)

	for n := 0; n < 10; n++ {
		for i := range nums {
			j := slices.Index(indices, i)
			indices = slices.Delete(indices, j, j+1)
			indices = slices.Insert(indices, int(modularIndex2((int64(j)+nums[i])%int64(len(indices)), int64(len(indices)))), i)
		}
	}
	zero := slices.Index(indices, slices.Index(nums, 0))

	v1 := nums[indices[modularIndex((zero+1000)%len(nums), len(nums))]]
	v2 := nums[indices[modularIndex((zero+2000)%len(nums), len(nums))]]
	v3 := nums[indices[modularIndex((zero+3000)%len(nums), len(nums))]]

	return v1 + v2 + v3
}

func Part2() {
	name := "Day #19 - part 2"

	common.TestOutputBig(name+" - input 2", 1623178306, Solve2(Input1))
	common.PrintOutputBig(name, Solve2(common.Readfile("./day20/input.txt")))
}
