package franc

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	five := NewFranc(5)
	if five.Times(2).Equals(NewFranc(10)) == false {
		t.Errorf("\nReturns if product of 5 and 2 equals 10, but return false")
	}
	if five.Times(3).Equals(NewFranc(15)) == false {
		t.Errorf("\nReturns if product of 5 and 3 equals 15, but return false")
	}
}

func TestEquality(t *testing.T) {
	res1 := NewFranc(5).Equals(NewFranc(5))
	if res1 != true {
		t.Errorf("\nReturn if 5 of Franc equals 5 of Franc\nExpected: %t, Actual: %t", true, res1)
	}
	res2 := NewFranc(5).Equals(NewFranc(6))
	if res2 == true {
		t.Errorf("\nReturn if 5 of Franc equals 6 of Franc\nExpected: %t, Actual: %t", false, res2)
	}
}
