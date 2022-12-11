package day11

import (
	"aoc/common"
	"regexp"
	"strconv"
	"strings"
)

type Monkey struct {
	index      int
	items      []int
	operation  string
	operand    string
	operandNum int
	test       int
	trueIndex  int
	falseIndex int
	inspected  int
}

func doOperation(operation, operand string, operandNum, div, old int) int {
	switch operation {
	case "+":
		switch operand {
		case "old":
			res := old % div
			return res + res
		default:
			return (old % div) + operandNum
		}
	case "*":
		switch operand {
		case "old":
			res := old % div
			return res * res
		default:
			return (old % div) * operandNum
		}
	}
	return -1
}

func (monkey Monkey) throw(monkeys *[]Monkey, ltm, dividing int) {
	for _, item := range monkey.items {
		(*monkeys)[monkey.index].inspected++
		newValue := doOperation(monkey.operation, monkey.operand, monkey.operandNum, ltm, item)
		if dividing > 1 {
			newValue = newValue / dividing
		}

		(*monkeys)[monkey.index].items = (*monkeys)[monkey.index].items[1:]
		if newValue%monkey.test == 0 {
			(*monkeys)[monkey.trueIndex].items = append((*monkeys)[monkey.trueIndex].items, newValue)
		} else {
			(*monkeys)[monkey.falseIndex].items = append((*monkeys)[monkey.falseIndex].items, newValue)
		}
	}
}

func parseMonkey(monkey string) Monkey {
	indexR, _ := regexp.Compile(`Monkey (\d+):`)
	indexStr := indexR.FindStringSubmatch(monkey)[1]
	opR, _ := regexp.Compile(`Operation: new = old (.) (\d+|old)`)
	op := opR.FindStringSubmatch(monkey)
	testR, _ := regexp.Compile(`Test: divisible by (\d+)`)
	testStr := testR.FindStringSubmatch(monkey)[1]
	trueR, _ := regexp.Compile(`If true: throw to monkey (\d+)`)
	trueStr := trueR.FindStringSubmatch(monkey)[1]
	falseR, _ := regexp.Compile(`If false: throw to monkey (\d+)`)
	falseStr := falseR.FindStringSubmatch(monkey)[1]

	itemLine := strings.Split(monkey, "\n")[1]
	re, _ := regexp.Compile(`\d+`)
	itemsStrs := re.FindAllString(itemLine, -1)
	items := []int{}

	for _, itemStr := range itemsStrs {
		item, _ := strconv.Atoi(itemStr)
		items = append(items, item)
	}

	index, _ := strconv.Atoi(indexStr)
	test, _ := strconv.Atoi(testStr)
	trueIndex, _ := strconv.Atoi(trueStr)
	falseIndex, _ := strconv.Atoi(falseStr)

	var operandNum int
	if op[2] == "old" {
		operandNum = 0
	} else {
		num, _ := strconv.Atoi(op[2])
		operandNum = num
	}

	return Monkey{index, items, op[1], op[2], operandNum, test, trueIndex, falseIndex, 0}
}

func product(monkeys []Monkey) int {
	p := 1
	for _, monkeys := range monkeys {
		p *= monkeys.test
	}
	return p
}

func iterate(numIterations int, monkeys []Monkey, dividing int) int64 {
	var max, max2 int64
	for i := 0; i < numIterations; i++ {
		for _, monkey := range monkeys {
			monkey.throw(&monkeys, product(monkeys), dividing)
		}
	}

	for _, monkey := range monkeys {
		ins := int64(monkey.inspected)
		if ins > max {
			max2 = max
			max = ins
		} else if ins > max2 {
			max2 = ins
		}
	}
	return max * max2
}

func Solve1(input string) int64 {
	monkeys := []Monkey{}
	for _, monkeyStr := range strings.Split(input, "\n\n") {
		monkeys = append(monkeys, parseMonkey(monkeyStr))
	}
	return iterate(20, monkeys, 3)
}

func Part1() {
	name := "Day #11 - part 1"

	common.TestOutputBig(name+" - input 1", 10605, Solve1(Input1))
	common.PrintOutputBig(name, Solve1(common.Readfile("./day11/input.txt")))
}
