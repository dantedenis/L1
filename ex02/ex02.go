package ex02

import "fmt"

func SqSlice(nums []int) {
	for _, n := range nums {
		// Инициализация и запуск горутины на каждый элемент из множества массива
		go func(n int) {
			fmt.Println(n * n)
		}(n)
	}
}
