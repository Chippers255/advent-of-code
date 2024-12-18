// generated by copilot for testing the sliceMath.go file
package tests

import (
	"testing"
	"year2024/internal/utils"
)

func TestSum(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	expectedInt := 15
	if result := utils.Sum(ints); result != expectedInt {
		t.Errorf("Sum(ints) = %d; want %d", result, expectedInt)
	}

	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	expectedFloat := 16.5
	if result := utils.Sum(floats); result != expectedFloat {
		t.Errorf("Sum(floats) = %f; want %f", result, expectedFloat)
	}
}

func TestAverage(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	expectedInt := 3.0
	if result := utils.Average(ints); result != expectedInt {
		t.Errorf("Average(ints) = %f; want %f", result, expectedInt)
	}

	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	expectedFloat := 3.3
	if result := utils.Average(floats); result != expectedFloat {
		t.Errorf("Average(floats) = %f; want %f", result, expectedFloat)
	}
}

func TestMax(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	expectedInt := 5
	if result := utils.Max(ints); result != expectedInt {
		t.Errorf("Max(ints) = %d; want %d", result, expectedInt)
	}

	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	expectedFloat := 5.5
	if result := utils.Max(floats); result != expectedFloat {
		t.Errorf("Max(floats) = %f; want %f", result, expectedFloat)
	}
}

func TestMin(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	expectedInt := 1
	if result := utils.Min(ints); result != expectedInt {
		t.Errorf("Min(ints) = %d; want %d", result, expectedInt)
	}

	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	expectedFloat := 1.1
	if result := utils.Min(floats); result != expectedFloat {
		t.Errorf("Min(floats) = %f; want %f", result, expectedFloat)
	}
}
