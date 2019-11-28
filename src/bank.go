package main

type Bank struct {
}

func (this *Bank) Reduce(source Expression, to string) IMoney {
	return source.Reduce(to)
}

func NewBank() *Bank {
	return &Bank{}
}
