package day03

import (
	"testing"

	"github.com/chippers255/advent-of-code/pkg/solver"
)

var samplePart1 = []byte(`xmul(2,4)%&mul[3,7]!@^mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)

var samplePart2 = []byte(`xmul(2,4)%&mul[3,7]!@^don't()mul(5,5)+mul(32,64]then(mul(11,8)do()mul(8,5))`)

func TestPart1(t *testing.T) {
	s := Solver{}
	got, err := s.Part1(solver.Context{Year: 2024, Day: 3, Input: samplePart1, Sample: true})
	if err != nil {
		t.Fatalf("Part1 error: %v", err)
	}
	if got != "161" {
		t.Fatalf("Part1 = %s, want %s", got, "161")
	}
}

func TestPart2(t *testing.T) {
	s := Solver{}
	got, err := s.Part2(solver.Context{Year: 2024, Day: 3, Input: samplePart2, Sample: true})
	if err != nil {
		t.Fatalf("Part2 error: %v", err)
	}
	if got != "48" {
		t.Fatalf("Part2 = %s, want %s", got, "48")
	}
}
