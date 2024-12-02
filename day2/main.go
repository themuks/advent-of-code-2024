package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reports := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		levels := make([]int, 0)
		for _, token := range tokens {
			level, _ := strconv.Atoi(token)
			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}

	count := 0

	for _, levels := range reports {
		if isSafe(levels) {
			count++
		}
	}

	fmt.Print(count)
}

func isSafe(levels []int) bool {
	const GAP = 3
	up := 0
	down := 0
	for i := 1; i < len(levels); i++ {
		d := levels[i] - levels[i-1]
		if d > GAP || d < -GAP {
			return false
		}
		if d > 0 {
			up++
		} else if d < 0 {
			down++
		} else {
			return false
		}
	}
	return !(up > 0 && down > 0 || up == 0 && down == 0)
}
