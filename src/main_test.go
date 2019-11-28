package main

import (
	"testing"

	"github.com/Kourin1996/TDD-go/src/bank"
	"github.com/Kourin1996/TDD-go/src/expression"
	"github.com/Kourin1996/TDD-go/src/money"
)

func TestSimpleAddition(t *testing.T) {
	five := money.NewDollar(5)
	sum := five.Plus(five)
	bank := bank.NewBank()
	reduced := bank.Reduce(sum, "USD")
	if isEqual := money.Equals(reduced, money.NewDollar(10)); isEqual != true {
		t.Errorf("\nReturn sum of 5 Dollar and 5 Dollar equals 10 Dollar\nExpected: %t, Actual: %t", true, false)
	}
}

func TestReduceSum(t *testing.T) {
	sum := expression.NewSum(money.NewDollar(3), money.NewDollar(4))
	bank := bank.NewBank()
	result := bank.Reduce(sum, "USD")
	if isEqual := money.Equals(money.NewDollar(7), result); isEqual != true {
		t.Errorf("\nReturn sum of 3 Dollar and 4 Dollar equals 7 Dollar\nExpected: %t, Actual: %t", true, false)
	}
}
