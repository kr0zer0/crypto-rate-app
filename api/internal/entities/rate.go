package entities

import "fmt"

type Rate struct {
	CurrencyPair
	Rate float64
}

func NewRate(pair CurrencyPair, rate float64) *Rate {
	return &Rate{
		CurrencyPair: pair,
		Rate:         rate,
	}
}

func (r *Rate) String() string {
	return fmt.Sprintf("%v - %v", r.CurrencyPair.String(), r.Rate)
}
