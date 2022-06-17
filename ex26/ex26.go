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
	fmt.Scan(&str)
	
	fmt.Println(uniqueSymbol(str))
}
