package ex11

func GetInter[T comparable](first, second []T) []T {
	result := make([]T, 0)
	myMap := make(map[T]bool)

	for _, t := range first {
		myMap[t] = true
	}

	for _, t := range second {
		if _, ok := myMap[t]; ok {
			result = append(result, t)
		}
	}

	return result
}
