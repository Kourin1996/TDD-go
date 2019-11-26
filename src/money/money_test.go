package money

import (
	"testing"
)

func TestMultiplication(t *testing.T) {
	if five := NewDollar(5); Equals(five.Times(3), NewDollar(10)) == true {
		t.Errorf("\nReturns if product of 5 and 3 equals 10, but return true")
	}
	if five := NewFranc(5); Equals(five.Times(2), NewFranc(10)) == false {
		t.Errorf("\nReturns if product of 5 and 2 equals 10, but return false")
	}
}

func TestEquality(t *testing.T) {
	if res := Equals(NewDollar(5), NewDollar(5)); res != true {
		t.Errorf("\nReturn if 5 of Dollar equals 5 of Dollar\nExpected: %t, Actual: %t", true, res)
	}
	if res := Equals(NewDollar(5), NewDollar(6)); res == true {
		t.Errorf("\nReturn if 5 of Dollar equals 6 of Dollar\nExpected: %t, Actual: %t", false, res)
	}
	if res := Equals(NewFranc(5), NewDollar(5)); res != false {
		t.Errorf("\nReturn if 5 of Franc equals 5 of Dollar\nExpected: %t, Actual: %t", false, res)
	}
}

func TestCurrency(t *testing.T) {
	if currency := NewDollar(1).GetCurrency(); currency != "USD" {
		t.Errorf("\nReturn currency of Dollar\nExpected: %s, Actual: %s", "USD", currency)
	}
	if currency := NewFranc(1).GetCurrency(); currency != "CHF" {
		t.Errorf("\nReturn currency of Dollar\nExpected: %s, Actual: %s", "CHF", currency)
	}
}

func TestSimpleAddition(t *testing.T) {
	sum := NewDollar(5).Plus(NewDollar(5))
	if isEqual := Equals(sum, NewDollar(10)); isEqual != true {
		t.Errorf("\nReturn sum of 5 Dollar and 5 Dollar equals 10 Dollar\nExpected: %t, Actual: %t", true, false)
	}
}
