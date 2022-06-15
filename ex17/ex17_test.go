package ex17

import (
	"sort"
	"testing"
)

func TestBinSearch(t *testing.T) {
	arr := []int{5, -1, 4, 16, -542, 2, 31, 12111, 1, 1, 2, 3, -1, 0, -1}
	sort.Ints(arr)

	target := 15
	index, err := BinSearch(arr, target)
	if err == nil {
		t.Error("Target found: ", arr[index], "excpected: ", target)
	}

	target = 1
	index, err = BinSearch(arr, target)
	if err != nil {
		t.Error(err)
	}
}
