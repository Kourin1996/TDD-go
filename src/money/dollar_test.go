package money

import (
	"testing"
)

func TestDollarMultiplication(t *testing.T) {
	five := NewDollar(5)
	if Equals(five.Times(2), NewDollar(10)) == false {
		t.Errorf("\nReturns if product of 5 and 2 equals 10, but return false")
	}
	if Equals(five.Times(3), NewDollar(10)) == true {
		t.Errorf("\nReturns if product of 5 and 3 equals 10, but return true")
	}
}

func TestDollarEquality(t *testing.T) {

	res1 := Equals(NewDollar(5), NewDollar(5))
	if res1 != true {
		t.Errorf("\nReturn if 5 of Dollar equals 5 of Dollar\nExpected: %t, Actual: %t", true, res1)
	}
	res2 := Equals(NewDollar(5), NewDollar(6))
	if res2 == true {
		t.Errorf("\nReturn if 5 of Dollar equals 6 of Dollar\nExpected: %t, Actual: %t", false, res2)
	}
}
