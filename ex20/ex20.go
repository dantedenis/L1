package ex20

import (
	"strings"
)

func ReverseWords(text string) string {
	words := strings.Split(text," ")
	for i,j := 0, len(words) - 1; i < j; i, j = i + 1, j - 1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}