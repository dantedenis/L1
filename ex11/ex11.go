package ex11

func GetInter[T comparable](first, second []T) []T {
	result := make([]T, 0)
	myMap := make(map[T]bool)

	// просто мапим каждое значение из первого слайса
	for _, t := range first {
		myMap[t] = true
	}

	// проверяем на совпадение значений в мапе и 2го слайса -> добавляем их в новый слайс
	for _, t := range second {
		if _, ok := myMap[t]; ok {
			result = append(result, t)
		}
	}

	return result
}

// Также можно решить за O(n^2): сравнив попарно все значения в 2х слайсах
