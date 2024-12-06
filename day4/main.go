package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "Part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		answer := part1(input)
		fmt.Println("Answer:", answer)
	} else {
		answer := part2(input)
		fmt.Println("Answer:", answer)
	}
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	emptyLine := strings.Repeat(".", len(lines))
	lines = append([]string{emptyLine}, lines...)
	lines = append(lines, emptyLine)
	for i := range lines {
		lines[i] = "." + lines[i] + "."
	}

	answer := 0

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if string(lines[i][j]) == "X" {
				answer += search(j, i, lines, "XMAS")
			}
		}
	}

	return 0
}

func search(x int, y int, lines []string, left string) int {
	fmt.Println(left)
	if string(lines[y][x]) == "." {
		return 0
	}

	if len(left) == 0 {
		return 1
	}

	if string(lines[x][y]) != string(left[0]) {
		return 0
	}

	/*
		00 10 20
		01    21
		02 12 22
	*/
	p00 := search(x-1, y-1, lines, left[1:])
	p10 := search(x, y-1, lines, left[1:])
	p20 := search(x+1, y-1, lines, left[1:])
	p21 := search(x+1, y, lines, left[1:])
	p22 := search(x+1, y+1, lines, left[1:])
	p12 := search(x, y+1, lines, left[1:])
	p02 := search(x-1, y+1, lines, left[1:])
	p01 := search(x-1, y, lines, left[1:])

	return p00 + p10 + p20 + p21 + p22 + p12 + p02 + p01
}

func part2(input string) int {
	return 0
}
