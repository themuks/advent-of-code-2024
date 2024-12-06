package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
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

type Rule struct {
	X int
	Y int
}

func part1(input string) int {
	rules, updateRows := parseInput(input)

	mappedRules := make(map[int][]int)
	for _, rule := range rules {
		_, ok := mappedRules[rule.X]
		if !ok {
			value := make([]int, 0)
			value = append(value, rule.Y)
			mappedRules[rule.X] = value
		} else {
			mappedRules[rule.X] = append(mappedRules[rule.X], rule.Y)
		}
	}

	answer := 0

	for _, updates := range updateRows {
		result := checkUpdateRow(updates, mappedRules)
		if result {
			answer += updates[len(updates)/2]
		}
	}

	return answer
}

func checkUpdateRow(row []int, rules map[int][]int) bool {
	for i, update := range row {
		for j := 0; j < i; j++ {
			if !checkComplience(row[j], rules[update]) {
				return false
			}
		}
	}

	return true
}

func checkComplience(update int, rules []int) bool {
	for _, rule := range rules {
		if rule == update {
			return false
		}
	}

	return true
}

func parseInput(input string) (rules []Rule, updateRows [][]int) {
	lines := strings.Split(input, "\n")

	var firstPart []string
	var secondPart []string
	for i, line := range lines {
		if line == "" {
			firstPart = lines[:i]
			secondPart = lines[i+1:]
			break
		}
	}

	for _, line := range firstPart {
		tokens := strings.Split(line, "|")
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		rules = append(rules, Rule{x, y})
	}

	for _, line := range secondPart {
		tokens := strings.Split(line, ",")
		updates := make([]int, 0)
		for _, token := range tokens {
			number, _ := strconv.Atoi(token)
			updates = append(updates, number)
		}
		updateRows = append(updateRows, updates)
	}

	return
}

func part2(input string) int {
	return 0
}
