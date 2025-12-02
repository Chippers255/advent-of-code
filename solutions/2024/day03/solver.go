package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/chippers255/advent-of-code/pkg/solver"
)

func init() {
	solver.Register(2024, 3, func() solver.Solver {
		return &Solver{}
	})
}

var (
	mulPattern = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	allPattern = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
)

// Solver contains the logic for Advent of Code 2024 Day 3.
type Solver struct{}

// Part1 sums every mul instruction found in the program output.
func (Solver) Part1(ctx solver.Context) (string, error) {
	input := strings.TrimSpace(string(ctx.Input))
	sum := 0
	for _, expr := range findMulExpressions(input) {
		arg1, arg2 := findExpressionArgs(expr)
		sum += arg1 * arg2
	}
	return fmt.Sprintf("%d", sum), nil
}

// Part2 respects do()/don't() controls while summing mul instructions.
func (Solver) Part2(ctx solver.Context) (string, error) {
	input := strings.TrimSpace(string(ctx.Input))
	sum := 0
	enabled := true

	for _, expr := range findAllExpressions(input) {
		switch expr {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				arg1, arg2 := findExpressionArgs(expr)
				sum += arg1 * arg2
			}
		}
	}

	return fmt.Sprintf("%d", sum), nil
}

func findMulExpressions(input string) []string {
	return mulPattern.FindAllString(input, -1)
}

func findAllExpressions(input string) []string {
	return allPattern.FindAllString(input, -1)
}

func findExpressionArgs(expr string) (int, int) {
	match := mulPattern.FindStringSubmatch(expr)
	if len(match) != 3 {
		return 0, 0
	}
	arg1, _ := strconv.Atoi(match[1])
	arg2, _ := strconv.Atoi(match[2])
	return arg1, arg2
}
