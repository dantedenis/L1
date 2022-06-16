package main

import (
	"fmt"
	"strings"
)

func uniqueSymbol(str string) bool {
	temp := []rune(strings.ToLower(str))
	mmap := make(map[rune]bool)
	
	for _, k := range temp {
		if _, ok := mmap[k]; ok {
			return false
		}
		mmap[k] = true
	}
	return true
}

func main() {
	var str string
	fmt.Println("Enter any strings")
	fmt.Scan(&str)
	
	fmt.Println(uniqueSymbol(str))
}