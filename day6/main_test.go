package main

import "testing"

var example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func Test_part1(t *testing.T) {
	want := 41
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
