package ex01

import (
	"fmt"
	"testing"
)

func TestHuman_String(t *testing.T) {
	h := NewHuman("Mark", "Twain", 15)
	if h.String() != "I'm Mark Twain, 15 y.o." {
		t.Error("String method is invalid")
	}
	a := Action{Human: *h, Skill: "programmer"}
	fmt.Println("1.", a.String())
	fmt.Println("2.", a)
}
