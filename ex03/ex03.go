package ex03

import (
	"sync"
)

func SumSqrt(nums []int) int {
	chSum := make(chan int, 1)
	//создание waitgroup
	wg := sync.WaitGroup{}
	chSum <- 0
	for _, n := range nums {
		wg.Add(1)
		go func(num int) {
			chSum <- num*num + <-chSum
			wg.Done()
		}(n)
	}
	wg.Wait()
	return <-chSum
}
