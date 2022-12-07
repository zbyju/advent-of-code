package day07

import (
	"aoc/common"
	"strings"
)

func findDirToDelete(dir Dir, spaceToFree int, freedSpace *int) (size int) {
	for _, v := range dir.files {
		size += v.size
	}

	for _, v := range dir.dirs {
		size += findDirToDelete(v, spaceToFree, freedSpace)
	}

	if size > spaceToFree && size < *freedSpace {
		*freedSpace = size
	}
	return size
}

func Solve2(input string) int {
	root := parseInput(strings.Split(input, "\n"))

	printDir(root, "/", 0)

	totalSpace := 70_000_000
	neededSpace := 30_000_000
	usedSpace := sumDirSizes(root, 0, nil)
	unusedSpace := totalSpace - usedSpace
	freedSpace := usedSpace

	findDirToDelete(root, neededSpace-unusedSpace, &freedSpace)
	return freedSpace
}

func Part2() {
	name := "Day #07 - part 2"

	common.TestOutput(name+" - input 1", 24933642, Solve2(Input1))
	common.PrintOutput(name, Solve2(common.Readfile("./day07/input.txt")))
}
