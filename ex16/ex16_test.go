package ex16

import (
	"testing"
	"sort"
)

func TestQuickSort(t *testing.T){
	arr1 := []int{5, -1, 4, 16, -542, 2, 31, 12111, 1, 1, 2, 3, -1, 0, -1}
	arr2 := make([]int, len(arr1))
	copy(arr2, arr1)
	
	QuickSort(arr1)
	sort.Ints(arr2)
	for i, _ := range arr1 {
		if arr1[i] != arr2[i]{
			t.Error("QuickSort dont work")
		}
	}
}