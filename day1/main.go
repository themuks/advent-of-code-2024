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

	freq := make(map[int]int)

	for i := range N {
		freq[right[i]] += 1
	}

	sum := 0

	for i := range N {
		count, ok := freq[left[i]]
		if !ok {
			continue
		}
		sum += count * left[i]
	}

	fmt.Print(sum)
}
