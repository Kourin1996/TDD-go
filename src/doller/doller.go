package doller

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{Amount: amount}
}

func (this *Dollar) times(multiplier int) {
	this.Amount = this.Amount * multiplier
}
