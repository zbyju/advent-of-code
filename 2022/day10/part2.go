package day10

import (
	"aoc/common"
	"fmt"
	"strconv"
	"strings"
)

type CPU2 struct {
	cycle  int
	X      int
	cycles []rune
	sprite []rune
}

func (cpu *CPU2) adjustSprite() {
	newSprite := []rune{}
	for i := 1; i <= 40; i++ {
		if i < cpu.X || i >= cpu.X+3 {
			newSprite = append(newSprite, '.')
		} else {
			newSprite = append(newSprite, '#')
		}
	}
	cpu.sprite = newSprite
}

func (cpu *CPU2) displayPixel() {
	cpu.cycles = append(cpu.cycles, cpu.sprite[(cpu.cycle+1)%len(cpu.sprite)])
	cpu.cycle++
}

func (cpu *CPU2) addx(num int) {
	cpu.displayPixel()
	cpu.X += num
	cpu.adjustSprite()
	cpu.displayPixel()
}

func (cpu *CPU2) noop() {
	cpu.displayPixel()
}

func runProgram2(lines []string) []rune {
	cpu := CPU2{
		0, 1, []rune{}, []rune("###....................................."),
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

func Solve2(input string) {
	cycles := runProgram2(strings.Split(input, "\n"))
	for i := 0; i < len(cycles); i++ {
		fmt.Print(string(cycles[i]))
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
}

func Part2() {
	name := "Day #10 - part 2"

	fmt.Println(name + " - input 1")
	Solve2(Input1)
	fmt.Println()

	fmt.Println(name + " - output is:")
	Solve2(common.Readfile("./day10/input.txt"))
}
