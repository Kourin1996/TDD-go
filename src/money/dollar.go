package money

type Dollar struct {
	*Money
}

func (this *Dollar) Times(multiplier int) IMoney {
	return &Dollar{Money: NewMoney(this.amount * multiplier)}
}

func (this *Dollar) GetCurrencyCode() string {
	return "Dollar"
}
