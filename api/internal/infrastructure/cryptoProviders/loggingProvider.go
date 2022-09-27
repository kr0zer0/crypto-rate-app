package crypto_providers

import (
	"api/internal/entities"
	"api/internal/usecases/usecases_contracts"
	"log"
	"reflect"
)

type LoggingCryptoProvider struct {
	cryptoProvider usecases_contracts.CryptoProvider
}

func NewLoggingCryptoProvider(cryptoProvider usecases_contracts.CryptoProvider) *LoggingCryptoProvider {
	return &LoggingCryptoProvider{cryptoProvider: cryptoProvider}
}

func (l *LoggingCryptoProvider) GetExchangeRate(currencyPair entities.CurrencyPair) (*entities.Rate, error) {
	rate, err := l.cryptoProvider.GetExchangeRate(currencyPair)
	if err != nil {
		log.Printf("%v - %v", l.getProviderName(), err)
		return nil, err
	}

	log.Printf("%v - %v", l.getProviderName(), rate.String())
	return rate, nil
}

func (l *LoggingCryptoProvider) getProviderName() string {
	return reflect.TypeOf(l.cryptoProvider).Elem().Name()
}
