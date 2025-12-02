# Advent of Code Harness

A reusable Go workspace for tackling Advent of Code puzzles across multiple years. It provides a CLI for running and scaffolding daily solutions plus a growing set of shared utilities.

## Requirements
- Go 1.21+

## Repository Layout
```
├── cmd/aoc/               # Cobra-based CLI
├── internal/templates/    # go:embed templates used by `aoc new`
├── inputs/<year>/         # Cached puzzle and sample inputs
├── pkg/
│   ├── input/             # Input loader and fetch helpers
│   ├── solver/            # Solver interfaces + registry
│   └── util/              # Shared helpers (parse, grid, mathx, strutil, ...)
├── solutions/<year>/dayXX # Per-day solver packages
└── scripts/               # (optional) helper scripts you add locally
```

## CLI Usage
All commands live under `./cmd/aoc` and can be run with `go run`.

Run a solution:
```
go run ./cmd/aoc run --year 2024 --day 1 --part both
```

Scaffold a new day (creates solver + tests + input files):
```
go run ./cmd/aoc new --year 2024 --day 5
```
Use `--force` to overwrite existing files.

## Inputs
Puzzle inputs live in `inputs/<year>/dayXX.txt` and optional samples in `inputs/<year>/dayXX-sample.txt`. The loader automatically selects the correct file based on the CLI flags.

## Writing Solvers
Each solver implements the `solver.Solver` interface and registers itself:
```go
func init() {
    solver.Register(2024, 1, func() solver.Solver { return &Solver{} })
}
```
The scaffolder generates a starter implementation that reads the input with helpers from `pkg/util/parse` so you can focus on puzzle logic.

## Utilities
Reusable helpers live under `pkg/util` (parsing lines/ints/grids, math helpers, string utilities, grid navigation, etc.). Add frequently used building blocks here and import them from any solver package.

## Testing
Each solver package comes with a `solver_test.go` to exercise sample inputs. Run all tests with `go test ./...`.
