package main

import "testing"

var example = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func Test_part1(t *testing.T) {
	want := 18
	if got := part1(example); got != want {
		t.Errorf("part1() = %v; want %v", got, want)
	}
}

var example2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func Test_part2(t *testing.T) {
	// want := 48
	// if got := part2(example2); got != want {
	// 	t.Errorf("part2() = %v; want %v", got, want)
	// }
}
