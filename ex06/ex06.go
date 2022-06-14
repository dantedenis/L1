package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ChanStop()
	ContextStop()
	WaitGroupStop()
	TimeoutStop()
	SignalStop()
}

func ChanStop() {
	fmt.Println("\tChannel Stop:")
	quit := make(chan bool)
	go func() {
		fmt.Println("Start goroutine")
		defer func() {
			fmt.Println("Stop goroutine")
		}()

		for {
			select {
			case <-quit:
				return
			}
		}
	}()

	fmt.Println("Send true flag")

	quit <- true
}

func SignalStop() {
	fmt.Println("\tSignal Stop:")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	go func() {
		for {
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}()
	<-sigChan
	fmt.Println("Stop goroutine")
}

func WaitGroupStop() {
	fmt.Println("\tWaitGroup Stop:")
	defer func() {
		fmt.Println("Stop goroutine")
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Working...")
		}
		wg.Done()
	}()
	wg.Wait()
}

func TimeoutStop() {
	fmt.Println("\tTimeout Stop:")
	go func() {
		for {
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Stop goroutine")
	}
}

func ContextStop() {
	fmt.Println("\tContext Stop:")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stop goroutine")
				wg.Done()
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
}
