package ex02

import "testing"

func TestSqSlice(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 15, 25, 45, 12, 13}

	SqSlice(nums)
}
