package money

import "fmt"

type IMoney interface {
	GetAmount() int
	GetCurrency() string
	Times(int) IMoney
}

func Equals(a, b IMoney) bool {
	return a.GetAmount() == b.GetAmount() && a.GetCurrency() == b.GetCurrency()
}

type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, currency string) *Money {
	return &Money{amount, currency}
}

func (this *Money) GetAmount() int {
	return this.amount
}

func (this *Money) GetCurrency() string {
	return this.currency
}

func (this *Money) Times(multiplier int) IMoney {
	return NewMoney(this.amount*multiplier, this.currency)
}

func (this *Money) ToString() string {
	return fmt.Sprintf("%d %s", this.amount, this.currency)
}

func NewDollar(amount int) *Dollar {
	return &Dollar{Money: NewMoney(amount, "USD")}
}

func NewFranc(amount int) *Franc {
	return &Franc{Money: NewMoney(amount, "CHF")}
}
