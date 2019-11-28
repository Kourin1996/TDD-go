package main

type Sum struct {
	Augend IMoney
	Addend IMoney
}

func NewSum(augend, addend IMoney) Expression {
	return &Sum{Augend: augend, Addend: addend}
}

func (this *Sum) Reduce(to string) IMoney {
	amount := this.Augend.GetAmount() + this.Addend.GetAmount()
	return NewMoney(amount, to)
}
