package money

import (
	"testing"
)

func TestDollarMultiplication(t *testing.T) {
	five := NewDollar(5)
	if five.Times(2).Equals(NewDollar(10)) == false {
		t.Errorf("\nReturns if product of 5 and 2 equals 10, but return false")
	}
	if five.Times(3).Equals(NewDollar(15)) == false {
		t.Errorf("\nReturns if product of 5 and 3 equals 15, but return false")
	}
}

func TestDollarEquality(t *testing.T) {
	res1 := NewDollar(5).Equals(NewDollar(5))
	if res1 != true {
		t.Errorf("\nReturn if 5 of Dollar equals 5 of Dollar\nExpected: %t, Actual: %t", true, res1)
	}
	res2 := NewDollar(5).Equals(NewDollar(6))
	if res2 == true {
		t.Errorf("\nReturn if 5 of Dollar equals 6 of Dollar\nExpected: %t, Actual: %t", false, res2)
	}
}
