package hw_1

import (
	"fmt"
	"testing"
)

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestAdd(t *testing.T) {
	results := [][]int{
		TwoSum([]int{2, 7, 11, 15}, 9),
		TwoSum([]int{3, 2, 4}, 6),
		TwoSum([]int{3, 3}, 6),
	}

	outputs := [][]int{
		{0, 1},
		{1, 2},
		{0, 1},
	}

	for index, res := range results {
		if !compareSlices(res, outputs[index]) {
			t.Error("failed")
		} else {
			fmt.Printf("test %d passed\n", index + 1)
		}
	}
}