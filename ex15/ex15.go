package main

import (
	"fmt"
	"strings"
)

/*
+ ко всему глобальная переменная, которую с легкостью можно перезаписать


var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}


Если тут подразумевается, что необходимо выполнить срез строки,
то "хвост" никуда не исчезнет и будет существовать все время пока работает с якобы срезом
т.к. ГК го запустится тогда, когда на эту область памяти никто не ссылается

для решения этого, необходимо явно выполнить копию данного участка слайса
*/

func main() {
	someFunc()
	someFuncNorm()
	fmt.Println(justString)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Printf("Start address: %p to next start str: %p\n", &v, &justString)
}

func someFuncNorm() {
	v := createHugeString(1 << 10)
	// явно клонируем/копируем срез строки
	justString := strings.Clone(v[:100])
	fmt.Printf("Start address: %p to next start str: %p\n", &v, &justString)
}

func createHugeString(num int) string {
	return strings.Repeat("q", num)
}
