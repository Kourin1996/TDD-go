package doller

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.times(2)
	if five.Amount != 10 {
		t.Errorf("\nReturn the product of 5 and 2\nExpected: %d, Actual: %d", 10, five.Amount)
	}
}
