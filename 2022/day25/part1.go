package day25

import (
	"aoc/common"
	"fmt"
	"math"
	"strings"
)

func SnafuToDec(snafu string) (res int) {
	for i, x := range snafu {
		y := 0
		switch x {
		case '=':
			y = -2
		case '-':
			y = -1
		case '0':
			y = 0
		case '1':
			y = 1
		case '2':
			y = 2
		}
		res += y * int(math.Pow(5.0, float64(len(snafu)-i-1)))
	}
	return res
}

func decToSnafu(num int) (res string) {
	for num > 0 {
		rem := num % 5
		num /= 5

		digit := ""
		switch rem {
		case 0:
			digit = "0"
		case 1:
			digit = "1"
		case 2:
			digit = "2"
		case 3:
			digit = "="
			num++
		case 4:
			digit = "-"
			num++
		}

		res = digit + res
	}
	return res
}

func parseInput(input string) []int {
	nums := []int{}
	for _, line := range strings.Split(input, "\n") {
		nums = append(nums, SnafuToDec(line))
	}
	return nums
}

func Solve1(input string) string {
	sum := 0
	for _, x := range parseInput(input) {
		sum += x
	}
	fmt.Println(sum)
	return decToSnafu(sum)
}

func Part1() {
	name := "Day #24 - part 1"

	common.TestOutputStr(name+" - input 1", "2=-1=0", Solve1(Input1))
	common.PrintOutputStr(name, Solve1(common.Readfile("./day25/input.txt")))
}
