package money

import (
	"testing"
)

func TestDollarFrancEquality(t *testing.T) {
	res1 := Equals(NewDollar(5), NewFranc(5))
	if res1 != false {
		t.Errorf("\nReturn if 5 of Dollar equals 5 of Franc\nExpected: %t, Actual: %t", false, res1)
	}
	res2 := Equals(NewFranc(5), NewDollar(6))
	if res2 != false {
		t.Errorf("\nReturn if 5 of Franc equals 6 of Dollar\nExpected: %t, Actual: %t", false, res2)
	}
}

func TestCurrency(t *testing.T) {
	if currency := NewDollar(1).GetCurrency(); currency != "USD" {
		t.Errorf("\nReturn currency of Dollar\nExpected: %t, Actual: %t", "USD", currency)
	}
	if currency := NewFranc(1).GetCurrency(); currency != "CHF" {
		t.Errorf("\nReturn currency of Dollar\nExpected: %t, Actual: %t", "CHF", currency)
	}
}
