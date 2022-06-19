package ex16

/*
Реализация по алгоритму https://en.wikipedia.org/wiki/Quicksort
 Основная суть:
	1. Выбрать опорный элемент
	2. Разбить массив: элементы больше опорного перемещаются перед ним, меньше - за ним
	3. Рекурсивно применять первые 2 шага к двум подмассивам слева и справа от опорного
*/

func QuickSort(arr []int) {
	if len(arr) > 1 {
		// Пунтк 1-2
		pivot := partition(arr)

		// Пунтк 3
		QuickSort(arr[:pivot])
		QuickSort(arr[pivot:])
	}
}

// Разбиение Ломуто
func partition(arr []int) int {
	last := len(arr) - 1
	pivot, i := arr[last], 0

	for j := 0; j < last; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[last] = arr[last], arr[i]
	return i
}

/*
	Библиотечный метод сортировки слайся с использованием компоратора

func SortSlice(arr []int) {
	sort.Slice(arr, func(i, j) bool {
		return arr[i] < arr[j]
	})
}

*/
