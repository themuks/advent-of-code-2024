package main

import "testing"

var example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func Test_part1(t *testing.T) {
	want := 3749
	if got := part1(example); got != want {
		t.Errorf("part1() = %v; want %v", got, want)
	}
}

func Test_part2(t *testing.T) {
	want := 0
	if got := part2(example); got != want {
		t.Errorf("part2() = %v; want %v", got, want)
	}
}
