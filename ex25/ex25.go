package main

import (
	"time"
	"context"
	"fmt"
)


func main() {
	fmt.Println(time.Now())
	SleepWithAfter(time.Second * 2)
	fmt.Println(time.Now())
	SleepWithLoop(time.Second * 2)
	fmt.Println(time.Now())
	SleepWithTick(time.Second * 2)
	fmt.Println(time.Now())
}

// Ожидание с помощью метода After, пока не вернет значени в канал через t время
func SleepWithAfter(t time.Duration) {
	<- time.After(t)
}

// Ожидание с помощью контекста с таймаутом t
func SleepWithCtx(t time.Duration) {
	end, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	<- end.Done()
}

// Цикл с проверкой прошло ли достаточно кол-ва времени t или нет
func SleepWithLoop(t time.Duration) {
	start := time.Now()
	
	for {
		end := start.Add(t)
		if time.Now().Sub(end) >= 0 {
			return
		}
	}
}

// слип с помощью тика, который пишет в канал после t времени
func SleepWithTick(t time.Duration) {
	ticks := time.Tick(t)
	<- ticks
}
