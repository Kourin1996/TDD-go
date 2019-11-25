package money

import (
	"testing"
)

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	if Equals(five.Times(2), NewFranc(10)) == false {
		t.Errorf("\nReturns if product of 5 and 2 equals 10, but return false")
	}
	if Equals(five.Times(3), NewFranc(10)) == true {
		t.Errorf("\nReturns if product of 5 and 3 equals 10, but return true")
	}
}

func TestFrancEquality(t *testing.T) {
	res1 := Equals(NewFranc(5), NewFranc(5))
	if res1 != true {
		t.Errorf("\nReturn if 5 of Franc equals 5 of Franc\nExpected: %t, Actual: %t", true, res1)
	}
	res2 := Equals(NewFranc(5), NewFranc(6))
	if res2 == true {
		t.Errorf("\nReturn if 5 of Franc equals 6 of Franc\nExpected: %t, Actual: %t", false, res2)
	}
}
