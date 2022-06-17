package ex10

func TemperatureGroup(temps []float64) map[int][]float64 {
	// инициализируем мапу
	result := make(map[int][]float64)
	for _, t := range temps {
		// делим на группы по сонованию 10
		group := int(t/10) * 10
		// сохраняем в мапе в необходимой группе
		result[group] = append(result[group], t)
	}
	return result
}
