package money

import (
	"testing"
)

func TestDollarFrancEquality(t *testing.T) {
	res1 := NewDollar(5).Equals(NewFranc(5))
	if res1 != false {
		t.Errorf("\nReturn if 5 of Dollar equals 5 of Franc\nExpected: %t, Actual: %t", false, res1)
	}
	res2 := NewFranc(5).Equals(NewDollar(6))
	if res2 != false {
		t.Errorf("\nReturn if 5 of Franc equals 6 of Dollar\nExpected: %t, Actual: %t", false, res2)
	}
}