package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var countWorks int
	fmt.Println("Enter count works:")
	fmt.Scan(&countWorks)
	chanMain := make(chan int)
	ctx, ctxClose := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer ctxClose()
	wg := sync.WaitGroup{}

	for countWorks > 0 {
		wg.Add(1)
		go func(ctx2 context.Context, id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Thread off", id)
					return
				case num := <-chanMain:
					fmt.Printf("Worker %d, read num: %d\n", id, num)
				}
			}
		}(ctx, countWorks)
		countWorks--
	}
	for {
		select {
		case <-ctx.Done():
			close(chanMain)
			wg.Wait()
			fmt.Println("Close main thread")
			return
		default:
			chanMain <- rand.Int()
			time.Sleep(500 * time.Millisecond)
		}
	}
}
