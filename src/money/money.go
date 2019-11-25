package money

type Money struct {
	amount int
}

func NewMoney(amount int) *Money {
	return &Money{amount: amount}
}

func (this *Money) Equals(that *Money) bool {
	return this.amount == that.amount
}
