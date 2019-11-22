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

func TestEquality(t *testing.T) {
	res1 := NewDollar(5).Equals(NewDollar(5))
	if res1 != true {
		t.Errorf("\nReturn if 5 of Dollar equals 5 of Dollar\nExpected: %t, Actual: %t", true, res1)
	}
	res2 := NewDollar(5).Equals(NewDollar(6))
	if res2 == true {
		t.Errorf("\nReturn if 5 of Dollar equals 6 of Dollar\nExpected: %t, Actual: %t", false, res2)
	}
}
