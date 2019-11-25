package money

type IMoney interface {
	GetAmount() int
	GetCurrencyCode() string
}

func Equals(a, b IMoney) bool {
	return a.GetAmount() == b.GetAmount() && a.GetCurrencyCode() == b.GetCurrencyCode()
}

type Money struct {
	amount int
}

func NewMoney(amount int) *Money {
	return &Money{amount}
}

func (this *Money) GetAmount() int {
	return this.amount
}
