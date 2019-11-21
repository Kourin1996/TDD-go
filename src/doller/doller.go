package doller

type Dollar struct {
	Amount int
}

func NewDollar(amount int) *Dollar {
	return &Dollar{Amount: 10}
}

func (this *Dollar) times(multiplier int) {

}
