package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var timeLife time.Duration
	fmt.Println("Enter time life for programm:")
	_, err := fmt.Scanln(&timeLife)
	if err != nil {
		log.Fatalln(err)
	}
	ctx, closer := context.WithTimeout(context.Background(), timeLife*time.Second)
	defer closer()
	chanMain := make(chan interface{})
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Thread off")
				return
			case t := <-chanMain:
				fmt.Println("Print data:", t)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			close(chanMain)
			wg.Wait()
			fmt.Println("Main routine finish")
			return
		case chanMain <- "message text":
			time.Sleep(1 * time.Second)
		}
	}
}
