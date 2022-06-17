package ex19

// достаточно бежать посимвольно и свапать их местами (метод 2 указателей)
func ReversString(str string) string {
	runes := []rune(str)
	for i,j := 0, len(runes) -1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
