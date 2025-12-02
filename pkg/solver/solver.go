package solver

// Solver describes a single AoC daily challenge implementation.
type Solver interface {
	Part1(Context) (string, error)
	Part2(Context) (string, error)
}

// Context bundles runtime details and puzzle input for a solver.
type Context struct {
	Year   int
	Day    int
	Input  []byte
	Sample bool
}

// Factory produces Solver implementations for registration.
type Factory func() Solver
