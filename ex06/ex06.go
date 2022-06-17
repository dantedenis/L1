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
		//инструкции перед завершением грутины
		defer func() {
			close(quit)
			fmt.Println("Stop goroutine")
		}()

		for {
			select {
			// Завершение горутины по принятию флагу в канал
			case <-quit:
				return
			}
		}
	}()

	fmt.Println("Send true flag")
	// Отправка сигнала на завершение горутины
	quit <- true
}

func SignalStop() {
	fmt.Println("\tSignal Stop:")
	sigChan := make(chan os.Signal, 1)
	// сам метод и горутина работаю, пока не будет принят сигнял прерывания
	signal.Notify(sigChan, syscall.SIGINT)
	go func() {
		for {
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}()
	// ожидание сигнала прерывания
	<-sigChan
	fmt.Println("Stop goroutine")
}

func WaitGroupStop() {
	fmt.Println("\tWaitGroup Stop:")
	defer func() {
		fmt.Println("Stop goroutine")
	}()
	wg := sync.WaitGroup{}
	// Группа горутин, завершение которых будем ожидать, в количестве 1шт
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("Working...")
		}
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
	// Завершение метода вместе с горутиной по таймеру
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Stop goroutine")
	}
}

func TimeoutStop2() {
	fmt.Println("\tTimeout Stop:")
	// Инициализация групп, для ожидания их завершения выполнения
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// Завершение горутины по таймеру
			select {
			case <-time.After(2 * time.Second):
				fmt.Println("Stop goroutine")
			default:
				fmt.Println("Working...")	
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Wait()
}

func ContextStop() {
	fmt.Println("\tContext Stop:")
	//Инициализация контекста с таймаутом, для завершения горутины по превышении времени выполнения
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func(){ 
			wg.Done()
			fmt.Println("Stop goroutine")
		}
		for {
			select {
			// Таймаут от контекста, завершение горутины
			case <-ctx.Done():
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
}
