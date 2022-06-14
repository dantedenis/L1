package ex03

import "testing"

func SumSqrt1(nums []int) int {
	sun := 0
	for _, n := range nums {
		sun += n * n
	}
	return sun
}

var (
	nums = []int{1, 2, 3, 4, 5, 6, 15, 2423, -1, 123}
)

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
