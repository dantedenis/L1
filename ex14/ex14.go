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

func DetectType(val interface{}) {
	fmt.Println("Type:", reflect.TypeOf(val))
}

func DetectTypePrintf(val interface{}) {
	fmt.Printf("_Type: %T\n", val)
}
