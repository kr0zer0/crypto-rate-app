package details

import (
	"api/internal/entities"
	"api/internal/usecases"
	"github.com/jellydator/ttlcache/v3"
	"time"
)

type CachedRateGetter struct {
	getRateUseCase *usecases.GetRateUseCase
	rateCache      *ttlcache.Cache[string, float64]
	cacheTTL       time.Duration
}

func NewCachedRateGetter(cryptoService *usecases.GetRateUseCase, cacheTTL time.Duration) *CachedRateGetter {
	rateCache := ttlcache.New[string, float64]()
	return &CachedRateGetter{
		getRateUseCase: cryptoService,
		rateCache:      rateCache,
		cacheTTL:       cacheTTL,
	}
}

func (c *CachedRateGetter) GetBtcUahRate() (*entities.Rate, error) {
	cachedItem := c.rateCache.Get("BTCUAH", ttlcache.WithDisableTouchOnHit[string, float64]())
	if cachedItem != nil && !cachedItem.IsExpired() {
		floatRate := cachedItem.Value()
		return entities.NewRate(entities.NewCurrencyPair(entities.BTC, entities.UAH), floatRate), nil
	}

	rate, err := c.getRateUseCase.GetBtcUahRate()
	if err != nil {
		return nil, err
	}

	c.rateCache.Set(rate.CurrencyPair.String(), rate.Rate, c.cacheTTL)

	return rate, nil
}
