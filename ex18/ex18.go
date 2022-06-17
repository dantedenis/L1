package main

import (
	"fmt"
)


//run with -race
func main(){
	runner := NewSyncCounter()
	for i := 0; i < 100; i++{
		// можно создать и отдельную группу горутин с ней и работать
		runner.Add(1)
		go func(){
			defer runner.Done()
			runner.Inc()
		}()
	}
	runner.Wait()
	
	fmt.Println(runner.GetCounter())
}
