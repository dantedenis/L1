package ex03

import (
	"sync"
	"sync/atomic"
)

func SumSqrt(nums []int) int64 {
	chSum := make(chan int64, 1)
	//создание waitgroup
	wg := sync.WaitGroup{}
	chSum <- 0
	for _, n := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			// забираем из канала значение и прибавляем к нему квадрат числа и тд
			chSum <- int64(num*num) + <-chSum
		}(n)
	}
	wg.Wait()
	return <-chSum
}

func SumSqrtAtomic(nums []int) (sum int64) {
	wg := sync.WaitGroup{}

	for _, n := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			// Атомарный доступ к значение сам и потокобезопасная запись в него результата
			atomic.AddInt64(&sum, int64(n*n))
		}(n)
	}
	wg.Wait()
	return sum
}
