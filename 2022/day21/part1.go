package day21

import (
	"aoc/common"
	"strconv"
	"strings"
)

type MonkeyOp struct {
	op    string
	left  string
	right string
}

func run(monkeys map[string]MonkeyOp, cache *map[string]int, curr string) int {
	if val, ok := (*cache)[curr]; ok {
		return val
	}
	leftVal := run(monkeys, cache, monkeys[curr].left)
	rightVal := run(monkeys, cache, monkeys[curr].right)
	val := 0
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

func Solve1(input string) int {
	cache := make(map[string]int)
	monkeys := make(map[string]MonkeyOp)
	for _, line := range strings.Split(input, "\n") {
		split1 := strings.Split(line, ": ")
		name := split1[0]
		if split1[1][0] >= '0' && split1[1][0] <= '9' {
			val, _ := strconv.Atoi(split1[1])
			cache[name] = val
		} else {
			split2 := strings.Split(split1[1], " ")
			monkeys[name] = MonkeyOp{split2[1], split2[0], split2[2]}
		}
	}
	return run(monkeys, &cache, "root")
}

func Part1() {
	name := "Day #21 - part 1"

	common.TestOutput(name+" - input 1", 152, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day21/input.txt")))
}
