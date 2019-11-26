package money

type Franc struct {
	*Money
}

func (this *Franc) Times(multiplier int) IMoney {
	return &Franc{Money: NewMoney(this.amount * multiplier)}
}

func (this *Franc) GetCurrencyCode() string {
	return "Franc"
}
