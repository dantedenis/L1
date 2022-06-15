package ex24

import (
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p Point) Distance(end Point) float64 {
	difX := math.Pow(p.x-end.x, 2)
	difY := math.Pow(p.y-end.y, 2)
	return math.Sqrt(difX + difY)
}

func Distance(first, second Point) float64 {
	return first.Distance(second)
}
