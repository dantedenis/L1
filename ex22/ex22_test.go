package ex22

import (
	"testing"
)

func TestNewBigInt(t *testing.T) {
	big := NewBigInt(1<<40, 1<<20)
	expect := "1099512676352"
	actual := big.Add()
	if actual.String() != expect {
		t.Error("Expected:", expect, ", Actual:", actual)
	}

	expect = "1099510579200"
	actual = big.Sub()
	if actual.String() != expect {
		t.Error("Expected:", expect, ", Actual:", actual)
	}

	expect = "1048576"
	actual = big.Div()
	if actual.String() != expect {
		t.Error("Expected:", expect, ", Actual:", actual)
	}

	expect = "1152921504606846976"
	actual = big.Mul()
	if actual.String() != expect {
		t.Error("Expected:", expect, ", Actual:", actual)
	}
}
