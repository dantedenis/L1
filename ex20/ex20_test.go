package ex20

import (
	"testing"
)

func TestReverse(t *testing.T){
	str := "snow dog sun"
	excp := "sun dog snow"
	if temp := ReverseWords(str); temp != excp {
		t.Error("Error reverse excpected:", excp, "actual:", temp)
	}
}