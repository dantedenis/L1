package ex23

import (
	"errors"
)

func RemoveElem[T any](arr []T, index int) ([]T, error) {
	if index < 0 || index > len(arr) {
		return nil, errors.New("index out of range")
	}
	return append(arr[:index], arr[index+1:]...), nil
}
