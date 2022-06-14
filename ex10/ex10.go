package ex10

func TemperatureGroup(temps []float64) map[int][]float64 {
	result := make(map[int][]float64)
	for _, t := range temps {
		group := int(t/10) * 10
		result[group] = append(result[group], t)
	}
	return result
}
