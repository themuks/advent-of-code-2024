package main

import "testing"

var example = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func Test_part1(t *testing.T) {
	want := 161
	if got := part1(example); got != want {
		t.Errorf("part1() = %v; want %v", got, want)
	}
}
