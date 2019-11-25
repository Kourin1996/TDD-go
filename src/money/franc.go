package money

type Franc struct {
	*Money
}

func NewFranc(amount int) *Franc {
	return &Franc{Money: NewMoney(amount)}
}

func (this *Franc) Times(multiplier int) *Franc {
	return &Franc{Money: NewMoney(this.amount * multiplier)}
}

func (this *Franc) GetCurrencyCode() string {
	return "Franc"
}
