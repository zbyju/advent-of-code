package day13

import (
	"aoc/common"
	"strconv"
	"strings"
)

type Packet struct {
	num    *int
	list   *[]Packet
	parent *Packet
}

type Pair struct {
	p1 Packet
	p2 Packet
}

func compare(p1 Packet, p2 Packet) int {
	if p1.num != nil {
		if p2.num != nil {
			return *p1.num - *p2.num
		}
		return compare(Packet{nil, &[]Packet{p1}, nil}, p2)
	} else if p2.num != nil {
		return compare(p1, Packet{nil, &[]Packet{p2}, nil})
	}
	minLen := len(*p1.list)
	if len(*p2.list) < minLen {
		minLen = len(*p2.list)
	}
	for i := 0; i < minLen; i++ {
		comp := compare((*p1.list)[i], (*p2.list)[i])
		if comp != 0 {
			return comp
		}
	}
	return len(*p1.list) - len(*p2.list)
}

func parsePacket(str string) Packet {
	p := Packet{nil, &[]Packet{}, nil}
	buf := ""
	cur := &p
	for _, v := range str[1:] {
		if v == ',' {
			val, _ := strconv.Atoi(buf)
			*cur.list = append(*cur.list, Packet{&val, nil, nil})
			buf = ""
		} else if v >= '0' && v <= '9' {
			buf += string(v)
		} else if v == '[' {
			buf = ""
			newPacket := Packet{nil, &[]Packet{}, cur}
			*cur.list = append(*cur.list, newPacket)
			cur = &newPacket
		} else if v == ']' {
			if len(buf) > 0 {
				val, _ := strconv.Atoi(buf)
				*cur.list = append(*cur.list, Packet{&val, nil, nil})
				buf = ""
			}
			if cur.parent == nil {
				return p
			}
			cur = cur.parent
		}
	}
	return p
}

func parseInput(pairsStr []string) []Pair {
	pairs := []Pair{}

	for _, p := range pairsStr {
		lines := strings.Split(p, "\n")
		p1 := parsePacket(lines[0])
		p2 := parsePacket(lines[1])
		pair := Pair{p1, p2}
		pairs = append(pairs, pair)
	}
	return pairs
}

func Solve1(input string) (sum int) {
	pairs := parseInput(strings.Split(input, "\n\n"))
	for i, p := range pairs {
		if compare(p.p1, p.p2) < 0 {
			sum += i + 1
		}
	}
	return sum
}

func Part1() {
	name := "Day #13 - part 1"

	common.TestOutput(name+" - input 1", 13, Solve1(Input1))
	common.PrintOutput(name, Solve1(common.Readfile("./day13/input.txt")))
}
