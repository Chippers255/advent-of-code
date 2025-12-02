package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/chippers255/advent-of-code/pkg/solver"
)

var (
	runYear   int
	runDay    int
	runPart   string
	runSample bool
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a specific Advent of Code solution",
	RunE: func(cmd *cobra.Command, args []string) error {
		if runYear == 0 || runDay == 0 {
			return errors.New("year and day must be provided")
		}

		part := strings.ToLower(runPart)
		if part == "" {
			part = "both"
		}

		inst, ctx, err := prepareSolver(runYear, runDay, runSample)
		if err != nil {
			return err
		}

		runPartFn := func(partLabel string, fn func(solver.Context) (string, error)) error {
			start := time.Now()
			answer, err := fn(ctx)
			duration := time.Since(start)
			if err != nil {
				return err
			}
			fmt.Printf("%s Part %s => %s (%s)\n", formatDay(runYear, runDay), partLabel, answer, duration)
			return nil
		}

		switch part {
		case "1", "p1", "one":
			return runPartFn("1", inst.Part1)
		case "2", "p2", "two":
			return runPartFn("2", inst.Part2)
		case "both", "all":
			if err := runPartFn("1", inst.Part1); err != nil {
				return err
			}
			return runPartFn("2", inst.Part2)
		default:
			return fmt.Errorf("unknown part %q", runPart)
		}
	},
}

func init() {
	runCmd.Flags().IntVarP(&runYear, "year", "y", 0, "AoC year to run")
	runCmd.Flags().IntVarP(&runDay, "day", "d", 0, "day of month to run")
	runCmd.Flags().StringVarP(&runPart, "part", "p", "both", "part to run: 1, 2, both")
	runCmd.Flags().BoolVar(&runSample, "sample", false, "use sample input if available")
	runCmd.MarkFlagRequired("year")
	runCmd.MarkFlagRequired("day")
}
