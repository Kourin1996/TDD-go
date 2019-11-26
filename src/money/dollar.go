package money

type Dollar struct {
	*Money
}

func (this *Dollar) Times(multiplier int) IMoney {
	return NewDollar(this.amount * multiplier)
}
