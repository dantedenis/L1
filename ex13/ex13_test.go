package ex13

import (
	"testing"
)

func TestSwap(t *testing.T) {
	first, second := 15, 2
	Swap(&first, &second)
	if first != 2 || second != 15 {
		t.Error("Error swap")
	}
}

func TestSwapSum(t *testing.T) {
	first, second := 15, 2
	SwapSum(&first, &second)
	if first != 2 || second != 15 {
		t.Error("Error swap")
	}
}
