package ex12

func MakeSet(words []string) []string {
	result := make([]string, 0)
	set := make(map[string]bool)

	for _, w := range words {
		set[w] = true
	}

	for k, _ := range set {
		if _, ok := set[k]; ok && k != "" {
			result = append(result, k)
		}
	}

	return result
}
