package main

import (
	"fmt"
	"strings"
)

func uniqueSymbol(str string) bool {
	temp := []rune(strings.ToLower(str))
	mmap := make(map[rune]bool)

	// мапим каждый символ, пока он не встретится в очередной раз
	for _, k := range temp {
		if _, ok := mmap[k]; ok {
			return false
		}
		mmap[k] = true
	}
	return true
}

// можно решить за n^2 делая полный перебор
func main() {
	var str string
	fmt.Println("Enter any strings")
	_, err := fmt.Scan(&str)
	if err != nil {
		return
	}

	fmt.Println(uniqueSymbol(str))
}
