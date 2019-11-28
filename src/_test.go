package money

import (
	"testing"

	"github.com/Kourin1996/TDD-go/src/bank"
	"github.com/Kourin1996/TDD-go/src/money"
)

func TestSimpleAddition(t *testing.T) {
	five := money.NewDollar(5)
	sum := five.Plus(five)
	bank := bank.NewBank()
	reduced := bank.Reduce(sum, "USD")
	if isEqual := Equals(reduced, NewDollar(10)); isEqual != true {
		t.Errorf("\nReturn sum of 5 Dollar and 5 Dollar equals 10 Dollar\nExpected: %t, Actual: %t", true, false)
	}
}
