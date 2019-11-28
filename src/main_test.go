package main

import (
	"testing"
)

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	if isEqual := Equals(reduced, NewDollar(10)); isEqual != true {
		t.Errorf("\nReturn sum of 5 Dollar and 5 Dollar equals 10 Dollar\nExpected: %t, Actual: %t", true, false)
	}
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	if isEqual := Equals(NewDollar(7), result); isEqual != true {
		t.Errorf("\nReturn sum of 3 Dollar and 4 Dollar equals 7 Dollar\nExpected: %t, Actual: %t", true, false)
	}
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.Reduce(NewDollar(1), "USD")
	if isEqual := Equals(NewDollar(1), result); isEqual != true {
		t.Errorf("\nReturn 1 Dollar equals 1 Dollar\nExpected: %t, Actual: %t", true, false)
	}
}
