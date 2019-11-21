package dollar

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	if product.Amount != 10 {
		t.Errorf("\nReturn the product of 5 and 2\nExpected: %d, Actual: %d", 10, product.Amount)
	}
	product = five.Times(3)
	if product.Amount != 15 {
		t.Errorf("\nReturn the product of 5 and 3\nExpected: %d, Actual: %d", 15, product.Amount)
	}
}
