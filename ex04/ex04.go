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
	_, err := fmt.Scan(&countWorks)
	if err != nil {
		fmt.Println(err)
		return
	}
	chanMain := make(chan int)
	// инициализация контекста для прерывания горутины по сигналу
	ctx, ctxClose := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer ctxClose()

	wg := sync.WaitGroup{}

	for countWorks > 0 {
		wg.Add(1)
		//запуск горутин с копией контекста на каждую
		go func(ctx2 context.Context, ctxClose context.CancelFunc, id int) {
			defer wg.Done()
			defer ctxClose()
			for {
				select {
				//отслеживание выполнения приема сигнала
				case <-ctx.Done():
					fmt.Println("Thread off", id)
					return
				case num := <-chanMain:
					fmt.Printf("Worker %d, read num: %d\n", id, num)
				}
			}
		}(ctx, ctxClose, countWorks)
		countWorks--
	}
	for {
		select {
		//ожидание завершения всех горутин или отправка в канал рандомного числа
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
