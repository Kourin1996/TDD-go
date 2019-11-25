package money

type Franc struct {
	Money
}

func NewFranc(amount int) *Franc {
	return &Franc{NewMoney(amount)}
}

func (this *Franc) Times(multiplier int) *Franc {
	return &Franc{NewMoney(this.amount * multiplier)}
}

func (this *Franc) Equals(that *Franc) bool {
	return this.amount == that.amount
}
