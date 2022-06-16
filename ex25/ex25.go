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


func SleepWithAfter(t time.Duration) {
	<- time.After(t)
}

func SleepWithCtx(t time.Duration) {
	end, cancel := context.WithTimeout(context.Background(), t)
	defer cancel()
	<- end.Done()
}

func SleepWithLoop(t time.Duration) {
	start := time.Now()
	
	for {
		end := start.Add(t)
		if time.Now().Sub(end) >= 0 {
			return
		}
	}
}

func SleepWithTick(t time.Duration) {
	ticks := time.Tick(t)
	<- ticks
}