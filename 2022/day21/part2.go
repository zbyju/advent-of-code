package day21

import (
	"aoc/common"
	"strconv"
	"strings"
)

func run2(monkeys map[string]MonkeyOp, cache *map[string]float64, curr string) float64 {
	if val, ok := (*cache)[curr]; ok {
		return val
	}

	leftVal := run2(monkeys, cache, monkeys[curr].left)
	rightVal := run2(monkeys, cache, monkeys[curr].right)

	if curr == "root" {
		if leftVal > rightVal {
			return -1
		} else if leftVal < rightVal {
			return 1
		}
		return 0
	}

	val := float64(0)
	switch monkeys[curr].op {
	case "+":
		val = leftVal + rightVal
	case "-":
		val = leftVal - rightVal
	case "*":
		val = leftVal * rightVal
	case "/":
		val = leftVal / rightVal
	}
	return val
}

func Solve2(input string) int64 {
	cache := make(map[string]float64)
	monkeys := make(map[string]MonkeyOp)
	for _, line := range strings.Split(input, "\n") {
		split1 := strings.Split(line, ": ")
		name := split1[0]
		if split1[1][0] >= '0' && split1[1][0] <= '9' {
			val, _ := strconv.Atoi(split1[1])
			cache[name] = float64(val)
		} else {
			split2 := strings.Split(split1[1], " ")
			monkeys[name] = MonkeyOp{split2[1], split2[0], split2[2]}
		}
	}

	lo, hi := int64(1), int64(1_000_000_000_000_000)
	for {
		mid := (hi-lo)/2 + lo
		cache["humn"] = float64(mid)

		cmp := run2(monkeys, &cache, "root")
		if cmp > 0 {
			hi = mid - 1
		} else if cmp < 0 {
			lo = mid + 1
		} else {
			return mid
		}
		if lo > hi {
			break
		}
	}

	lo, hi = int64(1), int64(1_000_000_000_000_000)
	for {
		mid := (hi-lo)/2 + lo
		cache["humn"] = float64(mid)

		cmp := run2(monkeys, &cache, "root")
		if cmp < 0 {
			hi = mid - 1
		} else if cmp > 0 {
			lo = mid + 1
		} else {
			return mid
		}
		if lo > hi {
			return -1
		}
	}
}

func Part2() {
	name := "Day #21 - part 2"

	common.TestOutputBig(name+" - input 1", 301, Solve2(Input1))
	common.PrintOutputBig(name, Solve2(common.Readfile("./day21/input.txt")))
}
