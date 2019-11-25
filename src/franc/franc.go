package franc

type Franc struct {
	amount int
}

func NewFranc(amount int) *Franc {
	return &Franc{amount: amount}
}

func (this *Franc) Times(multiplier int) *Franc {
	return &Franc{amount: this.amount * multiplier}
}

func (this *Franc) Equals(that *Franc) bool {
	return this.amount == that.amount
}
