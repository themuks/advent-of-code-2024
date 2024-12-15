package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
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

type Calibration struct {
	Total    int
	Operands []int
}

func part1(input string) int {
	calibrations := parseInput(input)

	answer := 0
	start := time.Now()

	for _, calibration := range calibrations {
		if search(calibration.Total, calibration.Operands, 0) {
			answer += calibration.Total
		}
	}

	fmt.Printf("Execution time: %v\n", time.Since(start))

	return answer
}

func search(total int, operands []int, current int) bool {
	return false
}

func parseInput(input string) (calibrations []Calibration) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		tokens := strings.Split(line, ": ")
		total, _ := strconv.Atoi(tokens[0])
		tokens = strings.Split(tokens[1], " ")
		operands := make([]int, len(tokens))
		for _, token := range tokens {
			operand, _ := strconv.Atoi(token)
			operands = append(operands, operand)
		}
		calibration := Calibration{total, operands}
		calibrations = append(calibrations, calibration)
	}

	return
}

func part2(input string) int {
	return 0
}
