package ex12

import (
	"fmt"
	"testing"
)

func TestMakeSet(t *testing.T) {
	words := []string{"cat", "cat", "cat", "dog", "tree", "pig", ""}
	fmt.Println(MakeSet(words))
}
