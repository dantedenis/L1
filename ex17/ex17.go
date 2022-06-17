package ex17

import (
	"errors"
)
// бинарный поиск эффективный для отсортированного массива данных
func BinSearch(arr []int, target int) (index int, err error) {
	if len(arr) < 1 {
		return 0, errors.New("Target not found")
	}
	
	// обращаемся к среднему элементу навсем участке поиска
	index = len(arr) / 2
	
	//рекурсивно будем вызывать поиск по участкам в массиве
	
	// если значение среднего больше цели то ищем в правой части подмассива
	if arr[index] > target {
		return BinSearch(arr[:index], target)
	// если меньше то ищем в левой
	} else if arr[index] < target {
		return BinSearch(arr[index+1:], target)
	}
	return index, nil
}
