package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	const N = 1000

	left := make([]int, 0)
	right := make([]int, 0)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "   ")
		parsedLeft, _ := strconv.Atoi(tokens[0])
		left = append(left, parsedLeft)
		parsedRight, _ := strconv.Atoi(tokens[1])
		right = append(right, parsedRight)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0

	for i := range N {
		if left[i] > right[i] {
			sum += left[i] - right[i]
		} else {
			sum += right[i] - left[i]
		}
	}

	fmt.Print(sum)
}
