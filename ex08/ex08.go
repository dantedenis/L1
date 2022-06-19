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
			// чисто проверка на конец файла/ввода (End of file)
			if err != io.EOF {
				log.Fatal(err)
			} else {
				fmt.Println("EOF exit")
			}
			return
		}
		SetBit(&number, target)
		PrintBytes(number)
	}

}

func SetBit(n *int64, pos int8) {
	if pos > 63 || pos < 0 {
		fmt.Println("Error: range out bit")
	} else {
		// Исключающее ИЛИ для сброса и установки бита в pos:  0010 ^ 0110 = 0100
		*n ^= 1 << pos
	}
}

func PrintBytes(n int64) {
	var i int8
	for i = 0; i < 64; i++ {
		// сдвиг вправо и побитовое И с текущим числом, для отображения состояния бита:  0010 0000 >> 1 = 0001 0000
		if (n>>i)&1 != 0 {
			fmt.Printf("1")
		} else {
			fmt.Printf("0")
		}
	}
	fmt.Println("\t", n)
}
