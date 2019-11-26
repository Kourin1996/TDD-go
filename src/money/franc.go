package money

type Franc struct {
	*Money
}

func (this *Franc) Times(multiplier int) IMoney {
	return NewFranc(this.amount * multiplier)
}
