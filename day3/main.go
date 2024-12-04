package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
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
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := r.FindAllStringSubmatch(input, -1)

	answer := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		answer += x * y
	}

	return answer
}

func part2(input string) int {
	return 0
}
