package main

import (
	"fmt"

	"github.com/chippers255/advent-of-code/pkg/input"
	"github.com/chippers255/advent-of-code/pkg/solver"
)

func prepareSolver(year, day int, sample bool) (solver.Solver, solver.Context, error) {
	loader := input.NewLoader("inputs")
	data, err := loader.Load(year, day, sample)
	if err != nil {
		return nil, solver.Context{}, err
	}

	inst, err := solver.Get(year, day)
	if err != nil {
		return nil, solver.Context{}, err
	}

	ctx := solver.Context{Year: year, Day: day, Input: data, Sample: sample}
	return inst, ctx, nil
}

func formatDay(year, day int) string {
	return fmt.Sprintf("%04d Day %02d", year, day)
}
