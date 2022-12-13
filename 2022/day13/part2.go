package day13

import (
	"aoc/common"
	"strings"
)

func parseInput2(pairsStr []string) []Packet {
	px := []Packet{}

	for _, p := range pairsStr {
		lines := strings.Split(p, "\n")
		p1 := parsePacket(lines[0])
		p2 := parsePacket(lines[1])
		px = append(px, p1)
		px = append(px, p2)
	}
	return px
}

func Solve2(input string) int {
	px := parseInput2(strings.Split(input, "\n\n"))
	v2, v6 := 2, 6
	p2 := Packet{nil, &[]Packet{{&v2, nil, nil}}, nil}
	p6 := Packet{nil, &[]Packet{{&v6, nil, nil}}, nil}
	i2 := 1
	i6 := 2
	for _, p := range px {
		if compare(p, p2) < 0 {
			i2++
			i6++
		} else if compare(p, p6) < 0 {
			i6++
		}
	}
	return i2 * i6
}

func Part2() {
	name := "Day #13 - part 2"

	common.TestOutput(name+" - input 2", 140, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day13/input.txt")))
}
