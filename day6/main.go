package main

import (
	_ "embed"
	"flag"
	"fmt"
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
	return 0
}

func part2(input string) int {
	return 0
}
