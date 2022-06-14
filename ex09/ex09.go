package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	numbers := make([]int, 0)
	for i := 0; i < rand.Int()%100; i++ {
		numbers = append(numbers, rand.Int()%1000)
	}
	chanOne, chanTwo := make(chan int), make(chan int)

	go func() {
		for _, num := range numbers {
			chanOne <- num
		}
		close(chanOne)
	}()

	go func() {
		for num := range chanOne {
			chanTwo <- num * num
		}
		close(chanTwo)
	}()

	for out := range chanTwo {
		fmt.Println(out)
	}
}
