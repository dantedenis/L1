package main

import (
	"math/rand"
	"sync"
)

//Start with "-race" flag
func main() {
	strs := []string{"one", "two", "three", "four", "five", "six", "second"}
	mmap := NewSafeMap()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mmap.Add(rand.Int()%7, strs[rand.Int()%7])
		}()
	}
	wg.Done()
}
