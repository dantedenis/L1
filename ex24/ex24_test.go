package ex24

import (
	"fmt"
	"testing"
)

func TestDist(t *testing.T) {
	p1 := NewPoint(2, 4)
	p2 := NewPoint(0, 0)

	fmt.Println(Distance(*p1, *p2))
}
