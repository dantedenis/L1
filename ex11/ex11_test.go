package ex11

import (
	"fmt"
	"testing"
)

func TestGetInter(t *testing.T) {
	first := []float64{1, 2, 3, 4, 5, 1, 1, 2, 3}
	second := []float64{6, 7, 8, 9, 10, 2, 1, 3}

	res := GetInter(first, second)
	fmt.Println(res)
}
