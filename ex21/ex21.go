package main

import (
	"fmt"
)

//Имеется интерфейс корты памяти, которую необходимо подключить к ПК
type CardMemory interface {
	Insert()
	Read() 
}

//Имплементация интерфейса карты памяти, но нет слота в ПК для коннекта
type CardMemoryM1 struct {
}

func (c CardMemoryM1) Insert() {
	fmt.Println("Succses connection")
}

func (c CardMemoryM1) Read() {
	fmt.Println("Succses read data")
}



//интерфейс юсб коннектора
type USB interface {
	ConnectPC()
}

//КардРидер имплеметирует интерфейс ЮСБ и является адаптером между ПК и картой памяти
type CardReader struct {
	card CardMemory
}

func NewCardReader(card CardMemory) *CardReader {
	return &CardReader{card: card}
}

func (c *CardReader) ConnectPC() {
	c.card.Insert()
	c.card.Read()
	fmt.Println("Succses connection PC")
}

func main() {
	reader := NewCardReader(CardMemoryM1{})
	reader.ConnectPC()
}