package money

type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{Money: NewMoney(amount)}
}

func (this *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{Money: NewMoney(this.amount * multiplier)}
}

func (this *Dollar) Equals(that *Dollar) bool {
	return this.Money.amount == that.Money.amount
}
