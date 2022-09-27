package entities

import "fmt"

type Currency string

const (
	BTC Currency = "BTC"
	UAH Currency = "UAH"
)

type CurrencyPair struct {
	base  Currency
	quote Currency
}

func NewCurrencyPair(base, quote Currency) CurrencyPair {
	return CurrencyPair{
		base:  base,
		quote: quote,
	}
}

func (p *CurrencyPair) GetBase() Currency {
	return p.base
}

func (p *CurrencyPair) GetQuote() Currency {
	return p.quote
}

func (p *CurrencyPair) String() string {
	return fmt.Sprintf("%v%v", p.quote, p.base)
}
