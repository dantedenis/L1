package ex19

import (
	"testing"
)

func TestReverse(t *testing.T){
	str := []string{"фбвгде", "absdefg", "WЁё╜Ї"}
	excp := []string{"едгвбф", "gfedsba", "Ї╜ёЁW"}
	for i, _:= range str {
		if temp := ReversString(str[i]); temp != excp[i] {
			t.Error("Error reverse excpected:", excp[i], "actual:", temp)
		}
	}
	
}