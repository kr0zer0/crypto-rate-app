package crypto_providers

import (
	"api/config"
	"api/internal/entities"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

type (
	BinanceProvider struct {
		APIUrl string
	}

	BinanceProviderCreator struct {
		cfg *config.Config
	}
)

func NewBinanceProviderCreator(cfg *config.Config) *BinanceProviderCreator {
	return &BinanceProviderCreator{cfg: cfg}
}

func (c *BinanceProviderCreator) CreateCryptoProvider() *BinanceProvider {
	return &BinanceProvider{
		APIUrl: c.cfg.CryptoProviders.Binance.URL,
	}
}

type binanceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (p *BinanceProvider) GetExchangeRate(currencyPair entities.CurrencyPair) (*entities.Rate, error) {
	response, err := p.makeAPIRequest(string(currencyPair.GetBase() + currencyPair.GetQuote()))
	if err != nil {
		return nil, err
	}

	var mappedResponse binanceResponse
	err = json.Unmarshal(response, &mappedResponse)
	if err != nil {
		return nil, err
	}

	price, err := strconv.ParseFloat(mappedResponse.Price, 64)
	if err != nil {
		return nil, fmt.Errorf("can't parse %v to float", price)
	}

	return entities.NewRate(currencyPair, price), nil
}

func (p *BinanceProvider) makeAPIRequest(symbol string) ([]byte, error) {
	client := resty.New()
	response, err := client.R().
		SetQueryParams(map[string]string{
			"symbol": symbol,
		}).
		Get(p.APIUrl)
	if err != nil {
		return nil, err
	}

	return response.Body(), nil
}
