package main

import (
	"fmt"
	"reflect"
)

func main() {
	DetectType(10)
	DetectType(true)
	DetectType("line")
	DetectType(make(chan int))
	DetectTypePrintf(10)
	DetectTypePrintf(true)
	DetectTypePrintf("line")
	DetectTypePrintf(make(chan int))
}

// определение типа с помощью пакета рефлект
func DetectType(val interface{}) {
	fmt.Println("Type:", reflect.TypeOf(val))
}

// можно просто принтф, но он так же использует рефлексию
func DetectTypePrintf(val interface{}) {
	fmt.Printf("_Type: %T\n", val)
}
