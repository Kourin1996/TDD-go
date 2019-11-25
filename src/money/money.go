package money

type Money struct {
	amount int
}

func NewMoney(amount int) Money {
	return Money{amount: amount}
}
