package ex03

import (
	"testing"
)

func SumSqrt1(nums []int) int {
	sun := 0
	for _, n := range nums {
		sun += n * n
	}
	return sun
}

var (
	nums     = []int{1, 2, 3, 4, 5, 6, 15, 2423, -1, 123}
	expected = int64(5886375)
)

func TestSumSqrt(t *testing.T) {
	a := SumSqrt(nums)
	if a != expected {
		t.Error("Expected: ", expected, " Actual: ", a)
	}
}

func TestSumSqrtAtomic(t *testing.T) {
	a := SumSqrtAtomic(nums)
	if a != expected {
		t.Error("Expected: ", expected, " Actual: ", a)
	}
}

func BenchmarkEx03_routing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSqrt(nums)
	}
}

func BenchmarkEx03_liner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSqrt1(nums)
	}
}
