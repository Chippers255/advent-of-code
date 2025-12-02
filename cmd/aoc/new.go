package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/chippers255/advent-of-code/internal/templates"
)

var (
	newYear  int
	newDay   int
	newForce bool
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Scaffold a new day's solution and input files",
	RunE: func(cmd *cobra.Command, args []string) error {
		if newYear == 0 || newDay == 0 {
			return errors.New("year and day must be provided")
		}

		pkgName := fmt.Sprintf("day%02d", newDay)
		solDir := filepath.Join("solutions", fmt.Sprintf("%d", newYear), pkgName)
		if err := os.MkdirAll(solDir, 0o755); err != nil {
			return err
		}

		solverPath := filepath.Join(solDir, "solver.go")
		testPath := filepath.Join(solDir, "solver_test.go")

		solverContent, err := templates.RenderDaySolver(newYear, newDay)
		if err != nil {
			return err
		}
		if err := writeFileMaybe(solverPath, solverContent, newForce); err != nil {
			return err
		}

		testContent, err := templates.RenderDayTest(newYear, newDay)
		if err != nil {
			return err
		}
		if err := writeFileMaybe(testPath, testContent, newForce); err != nil {
			return err
		}

		inputDir := filepath.Join("inputs", fmt.Sprintf("%d", newYear))
		if err := os.MkdirAll(inputDir, 0o755); err != nil {
			return err
		}
		inputPath := filepath.Join(inputDir, fmt.Sprintf("day%02d.txt", newDay))
		if err := writeFileMaybe(inputPath, []byte("# paste puzzle input here\n"), false); err != nil {
			return err
		}
		samplePath := filepath.Join(inputDir, fmt.Sprintf("day%02d-sample.txt", newDay))
		if err := writeFileMaybe(samplePath, []byte("# optional sample input\n"), false); err != nil {
			return err
		}

		fmt.Printf("Scaffolded %s in %s\n", formatDay(newYear, newDay), solDir)
		return nil
	},
}

func init() {
	newCmd.Flags().IntVarP(&newYear, "year", "y", 0, "AoC year to scaffold")
	newCmd.Flags().IntVarP(&newDay, "day", "d", 0, "day of month to scaffold")
	newCmd.Flags().BoolVar(&newForce, "force", false, "overwrite existing files")
	newCmd.MarkFlagRequired("year")
	newCmd.MarkFlagRequired("day")
}

func writeFileMaybe(path string, data []byte, force bool) error {
	if _, err := os.Stat(path); err == nil && !force {
		return fmt.Errorf("%s already exists (use --force to overwrite)", path)
	} else if err == nil && force {
		return os.WriteFile(path, data, 0o644)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return err
	}
	return nil
}
