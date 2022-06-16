package ex02

import (
	"fmt"
	"sync"
)

//Порядок вывода не гарантируется
func SqSlice(nums []int) {
	//Коллекция для горутин, необходимо что бы контролировать завершение всех горутин
	wg := sync.WaitGroup{}
	for _, n := range nums {
		wg.Add(1)
		// Инициализация и запуск горутины на каждый элемент из множества массива
		go func(n int) {
			defer wg.Done()
			fmt.Println(n * n)
		}(n)
	}
	//Ожидание завершения всех горутин
	wg.Wait()
}
