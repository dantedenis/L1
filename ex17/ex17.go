package ex17

import (
	"errors"
)
//
func BinSearch(arr []int, target int) (index int, err error) {
	if len(arr) < 1 {
		return 0, errors.New("Target not found")
	}
	
	index = len(arr) / 2
	
	if arr[index] > target {
		return BinSearch(arr[:index], target)
	} else if arr[index] < target {
		return BinSearch(arr[index+1:], target)
	}
	return index, nil
}