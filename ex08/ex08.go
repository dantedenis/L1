package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	var number int64
	var target int8
	for {

		_, err := fmt.Scan(&target)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			} else {
				fmt.Println("EOF exit")
			}
			return
		}
		SetBit(&number, target)
		PrintBytes(number)
		fmt.Println(number)
	}

}

func SetBit(n *int64, pos int8) {
	if pos > 63 || pos < 0 {
		fmt.Println("Error: range out bit")
	} else {
		*n ^= 1 << pos
	}
}

func PrintBytes(n int64) {
	var i int8
	for i = 0; i < 64; i++ {
		if (n & (1 << i)) != 0 {
			fmt.Printf("1")
		} else {
			fmt.Printf("0")
		}
	}
	fmt.Println()
}
