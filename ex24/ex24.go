package ex24

import (
	"math"
)

// структура поинт с неэкспортируемыми полями
type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// метод структуры для нахождения дистанции
func (p Point) Distance(end Point) float64 {
	difX := math.Pow(p.x-end.x, 2)
	difY := math.Pow(p.y-end.y, 2)
	return math.Sqrt(difX + difY)
}

//экспортируемый метод от 2 поинтов
func Distance(first, second Point) float64 {
	return first.Distance(second)
}
