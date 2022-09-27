package crypto_providers

import (
	"api/config"
	"api/internal/entities"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type (
	CoinAPIProvider struct {
		HeaderName string
		APIKey     string
		APIUrl     string
	}

	CoinAPIProviderCreator struct {
		cfg *config.Config
	}
)

func NewCoinAPIProviderCreator(cfg *config.Config) *CoinAPIProviderCreator {
	return &CoinAPIProviderCreator{cfg: cfg}
}

func (c *CoinAPIProviderCreator) CreateCryptoProvider() *CoinAPIProvider {
	return &CoinAPIProvider{
		HeaderName: c.cfg.CryptoProviders.CoinAPI.HeaderName,
		APIKey:     c.cfg.CryptoProviders.CoinAPI.APIKey,
		APIUrl:     c.cfg.CryptoProviders.CoinAPI.URL,
	}
}

type coinAPIResponse struct {
	Rate float64 `json:"rate"`
}

func (p *CoinAPIProvider) GetExchangeRate(currencyPair entities.CurrencyPair) (*entities.Rate, error) {
	response, err := p.makeAPIRequest(string(currencyPair.GetBase()), string(currencyPair.GetQuote()))
	if err != nil {
		return nil, err
	}

	var mappedResponse coinAPIResponse
	err = json.Unmarshal(response, &mappedResponse)
	if err != nil {
		return nil, err
	}

	return entities.NewRate(currencyPair, mappedResponse.Rate), nil
}

func (p *CoinAPIProvider) makeAPIRequest(baseCurrency, quoteCurrency string) ([]byte, error) {
	client := resty.New()
	response, err := client.R().
		SetPathParams(map[string]string{
			"base":  baseCurrency,
			"quote": quoteCurrency,
		}).
		SetHeader(p.HeaderName, p.APIKey).
		Get(p.APIUrl)
	if err != nil {
		return nil, err
	}

	return response.Body(), nil
}
