package day02

import (
	"testing"

	"github.com/chippers255/advent-of-code/pkg/solver"
)

var sampleInput = []byte(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`)

func TestPart1(t *testing.T) {
	s := Solver{}
	got, err := s.Part1(solver.Context{Year: 2024, Day: 2, Input: sampleInput, Sample: true})
	if err != nil {
		t.Fatalf("Part1 error: %v", err)
	}
	if got != "2" {
		t.Fatalf("Part1 = %s, want %s", got, "2")
	}
}

func TestPart2(t *testing.T) {
	s := Solver{}
	got, err := s.Part2(solver.Context{Year: 2024, Day: 2, Input: sampleInput, Sample: true})
	if err != nil {
		t.Fatalf("Part2 error: %v", err)
	}
	if got != "4" {
		t.Fatalf("Part2 = %s, want %s", got, "4")
	}
}
