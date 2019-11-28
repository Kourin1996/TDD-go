package bank

import (
	"github.com/Kourin1996/TDD-go/src/expression"
	"github.com/Kourin1996/TDD-go/src/money"
)

type Bank struct {
}

func (this *Bank) Reduce(source expression.Expression, to string) money.IMoney {
	return money.NewDollar(10)
}

func NewBank() *Bank {
	return &Bank{}
}
