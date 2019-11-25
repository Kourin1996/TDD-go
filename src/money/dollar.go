package money

type Dollar struct {
	Money
}

func NewDollar(amount int) *Dollar {
	return &Dollar{NewMoney(amount)}
}

func (this *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{NewMoney(this.amount * multiplier)}
}

func (this *Dollar) Equals(that *Dollar) bool {
	return this.amount == that.amount
}
