package day01

import (
	"fmt"
	"sort"

	"github.com/chippers255/advent-of-code/pkg/input"
	"github.com/chippers255/advent-of-code/pkg/solver"
	"github.com/chippers255/advent-of-code/pkg/util/mathx"
)

func init() {
	solver.Register(2024, 1, func() solver.Solver {
		return &Solver{}
	})
}

// Solver contains the logic for Advent of Code 2024 Day 1.
type Solver struct{}

func parseLists(ctx solver.Context) ([]int, []int, error) {
	left, right, err := input.ParseIntPairs(ctx.Input)
	if err != nil {
		return nil, nil, err
	}
	if len(left) != len(right) {
		return nil, nil, fmt.Errorf("mismatched pair counts: %d vs %d", len(left), len(right))
	}
	return left, right, nil
}

// Part1 solves the first part of the puzzle.
func (Solver) Part1(ctx solver.Context) (string, error) {
	left, right, err := parseLists(ctx)
	if err != nil {
		return "", err
	}

	sort.Ints(left)
	sort.Ints(right)

	diffs := make([]int, len(left))
	for i := range left {
		diffs[i] = mathx.AbsDiff(left[i], right[i])
	}

	return fmt.Sprintf("%d", mathx.SumInts(diffs)), nil
}

// Part2 solves the second part of the puzzle.
func (Solver) Part2(ctx solver.Context) (string, error) {
	left, right, err := parseLists(ctx)
	if err != nil {
		return "", err
	}

	counts := make(map[int]int, len(right))
	for _, value := range right {
		counts[value]++
	}

	total := 0
	for _, value := range left {
		if c := counts[value]; c > 0 {
			total += value * c
		}
	}

	return fmt.Sprintf("%d", total), nil
}
