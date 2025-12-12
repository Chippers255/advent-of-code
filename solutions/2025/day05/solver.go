package day05

import (
	"fmt"
	"sort"

	"github.com/chippers255/advent-of-code/pkg/solver"
	"github.com/chippers255/advent-of-code/pkg/util/parse"
)

func init() {
	solver.Register(2025, 5, func() solver.Solver {
		return &Solver{}
	})
}

// Solver contains the logic for Advent of Code 2025 Day 5.
type Solver struct{}

// Part1 solves the first part of the puzzle.
func (Solver) Part1(ctx solver.Context) (string, error) {
	ranges, ints, err := parse.RangeLinesAndInts(ctx.Input)
	if err != nil {
		return "", err
	}

	answer := 0
	for _, i := range ints {
		for _, r := range ranges {
			if i >= r[0] && i <= r[1] {
				answer++
				break
			}
		}
	}
	return fmt.Sprintf("%d", answer), nil
}

func mergeRanges(ranges [][]int) [][]int {
	if len(ranges) == 0 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := [][]int{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := merged[len(merged)-1]
		current := ranges[i]

		// If current range overlaps or is adjacent to the last merged range
		if current[0] <= last[1]+1 {
			// Merge: extend the end if current range extends further
			if current[1] > last[1] {
				merged[len(merged)-1][1] = current[1]
			}
		} else {
			// No overlap, add as new range
			merged = append(merged, current)
		}
	}

	return merged
}

// Part2 solves the second part of the puzzle.
func (Solver) Part2(ctx solver.Context) (string, error) {
	ranges, _, err := parse.RangeLinesAndInts(ctx.Input)
	if err != nil {
		return "", err
	}

	merged := mergeRanges(ranges)
	answer := 0
	for _, r := range merged {
		// Count IDs in range: end - start + 1
		answer += r[1] - r[0] + 1
	}

	return fmt.Sprintf("%d", answer), nil
}
