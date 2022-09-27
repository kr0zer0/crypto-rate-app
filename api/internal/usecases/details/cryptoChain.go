package details

import (
	"api/internal/entities"
	"api/internal/usecases/usecases_contracts"
)

type CryptoChain interface {
	usecases_contracts.CryptoProvider
	SetNext(CryptoChain)
}

type BaseCryptoChain struct {
	next     CryptoChain
	provider usecases_contracts.CryptoProvider
}

func NewBaseCryptoChain(provider usecases_contracts.CryptoProvider) CryptoChain {
	return &BaseCryptoChain{provider: provider}
}

func (c *BaseCryptoChain) SetNext(next CryptoChain) {
	c.next = next
}

func (c *BaseCryptoChain) GetExchangeRate(currencyPair entities.CurrencyPair) (*entities.Rate, error) {
	rate, err := c.provider.GetExchangeRate(currencyPair)
	if err != nil {
		nextChain := c.next
		if nextChain == nil {
			return nil, err
		}

		rate, err = nextChain.GetExchangeRate(currencyPair)
	}

	return rate, err
}
