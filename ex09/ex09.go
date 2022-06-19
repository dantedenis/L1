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
	// Для конвейера каналов будем использовать 2 канала.
	chanOne, chanTwo := make(chan int), make(chan int)

	// эта горутина будет бежать по массиву и отправлять значения в канал и закрывает его в конечном итоге
	go func() {
		defer close(chanOne)
		for _, num := range numbers {
			chanOne <- num
		}
	}()

	// эта горутина читает из канала, возводит в квадрат и отправляет след канал
	go func() {
		defer close(chanTwo)
		for num := range chanOne {
			chanTwo <- num * num
		}
	}()

	// основная горутина читает из второго канала и выводит на stdout
	for out := range chanTwo {
		fmt.Println(out)
	}
}
