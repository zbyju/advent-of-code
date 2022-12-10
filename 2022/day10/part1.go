package day10

import (
	"aoc/common"
	"strconv"
	"strings"
)

type CPU struct {
	X      int
	cycles []int
}

func (cpu *CPU) addx(num int) {
	cpu.cycles = append(cpu.cycles, cpu.X)
	cpu.X += num
	cpu.cycles = append(cpu.cycles, cpu.X)
}

func (cpu *CPU) noop() {
	cpu.cycles = append(cpu.cycles, cpu.X)
}

func runProgram(lines []string) []int {
	cpu := CPU{
		1, []int{1},
	}
	for _, line := range lines {
		split := strings.Split(line, " ")
		op, args := split[0], split[1:]
		switch op {
		case "addx":
			num, _ := strconv.Atoi(args[0])
			cpu.addx(num)
		default:
			cpu.noop()
		}
	}
	return cpu.cycles
}

func Solve1(input string) (sum int) {
	cycles := runProgram(strings.Split(input, "\n"))
	for i := 19; i < len(cycles); i += 40 {
		sum += (i + 1) * cycles[i]
	}
	return sum
}

func Part1() {
	name := "Day #10 - part 1"

	common.TestOutput(name+" - input 1", 13140, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day10/input.txt")))
}
