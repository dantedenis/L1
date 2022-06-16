package ex01

import "fmt"

type Human struct {
	age       int
	firstName string
	lastName  string
}

//Имплементирую интерфейс Stringer
func (h *Human) String() string {
	return fmt.Sprintf("I'm %s %s, %d y.o.", h.firstName, h.lastName, h.age)
}

func NewHuman(firstName, lastName string, age int) *Human {
	return &Human{firstName: firstName, lastName: lastName, age: age}
}

// Action : Встраивание(композиция) структуры хуман в структуру акшион
type Action struct {
	Human
	Skill string
}

// Action1 : Встраивание указателя на структуру
type Action1 struct {
	*Human
	Skill string
}
