package mathx

// MinInt returns the smaller of a or b.
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the larger of a or b.
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// AbsInt returns the absolute value of n.
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// AbsDiff returns the absolute difference between a and b.
func AbsDiff(a, b int) int {
	return AbsInt(a - b)
}

// SumInts returns the sum of the provided integers.
func SumInts(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// FloorDiv performs integer division rounding toward negative infinity.
func FloorDiv(a, b int) int {
	if b == 0 {
		panic("mathx.FloorDiv: division by zero")
	}
	q := a / b
	r := a % b
	if (r != 0) && ((r > 0) != (b > 0)) {
		q--
	}
	return q
}
