package ex10

import (
	"fmt"
	"testing"
)

func TestTemperatureGroup(t *testing.T) {
	temps := []float64{-25.0, 2, -27.0, 34.5, 19.0, 11, 5, 15.6, 10.1, -1, 0}
	result := TemperatureGroup(temps)

	for k, v := range result {
		fmt.Println(k, v)
	}
}
