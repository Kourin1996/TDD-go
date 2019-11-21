package doller

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{Amount: amount}
}

func (this *Dollar) Times(multiplier int) *Dollar {
	return &Dollar{Amount: this.Amount * multiplier}
}
