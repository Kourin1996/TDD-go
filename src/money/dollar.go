package money

type Dollar struct {
	amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{amount: amount}
}

func (this *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{amount: this.amount * multiplier}
}

func (this *Dollar) Equals(that *Dollar) bool {
	return this.amount == that.amount
}
