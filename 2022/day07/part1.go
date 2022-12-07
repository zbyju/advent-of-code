package day07

import (
	"aoc/common"
	"strconv"
	"strings"
)

type Dir struct {
	parent *Dir
	files  map[string]File
	dirs   map[string]Dir
}

type File struct {
	size int
}

type FS struct {
	currentDir *Dir
	root       *Dir
}

func doCd(path string, fs *FS) {
	switch path {
	case "/":
		fs.currentDir = fs.root
	case "..":
		fs.currentDir = fs.currentDir.parent
	default:
		dir := fs.currentDir.dirs[path]
		fs.currentDir = &dir
	}
}

func doLs(output []string, fs *FS) {
	for _, out := range output {
		isDir, name, size := parseOutput(out)
		if _, ok := fs.currentDir.dirs[name]; isDir && !ok {
			fs.currentDir.dirs[name] = Dir{fs.currentDir, make(map[string]File), make(map[string]Dir)}
		} else if _, ok := fs.currentDir.files[name]; !isDir && !ok {
			fs.currentDir.files[name] = File{size}
		}
	}
}

func parseOutput(out string) (bool, string, int) {
	split := strings.Split(out, " ")

	if split[0] == "dir" {
		return true, split[1], 0
	}
	size, _ := strconv.Atoi(split[0])
	return false, split[1], size
}

func getOutput(lines []string, index int) (output []string) {
	for i := index; i < len(lines); i++ {
		if lines[i][0] == '$' {
			return output
		}
		output = append(output, lines[i])
	}
	return output
}

func parseInput(lines []string) Dir {
	root := Dir{nil, make(map[string]File), make(map[string]Dir)}
	fs := FS{&root, &root}
	for i, line := range lines {
		if strings.HasPrefix(line, "$ ls") {
			output := getOutput(lines, i+1)
			doLs(output, &fs)
		} else if strings.HasPrefix(line, "$ cd") {
			split := strings.Split(line, " ")
			doCd(split[2], &fs)
		}
	}
	return *fs.root
}

func sumDirSizes(dir Dir, maxSize int, sum *int64) (size int) {
	for _, v := range dir.files {
		size += v.size
	}

	for _, v := range dir.dirs {
		size += sumDirSizes(v, maxSize, sum)
	}

	if sum != nil && size <= maxSize {
		*sum += int64(size)
	}
	return size
}

func Solve1(input string) (sum int64) {
	root := parseInput(strings.Split(input, "\n"))

	sumDirSizes(root, 100_000, &sum)

	return sum
}

func Part1() {
	name := "Day #07 - part 1"

	common.TestOutputBig(name+" - input 1", 95437, Solve1(Input1))
	common.PrintOutputBig(name, Solve1(common.Readfile("./day07/input.txt")))
}
