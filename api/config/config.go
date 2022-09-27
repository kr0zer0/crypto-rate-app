package config

import (
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		Port string `env-required:"true" yaml:"port"`
	} `yaml:"app"`
	CryptoProviders struct {
		CryptoProvider string `env-required:"true" env:"CRYPTO_CURRENCY_PROVIDER"`
		CoinMarketCap  struct {
			URL        string `env-required:"true" yaml:"url"`
			HeaderName string `env-required:"true" yaml:"headerName"`
			APIKey     string `env-required:"true" env:"COINMARKETCAP_API_KEY"`
		} `yaml:"coinMarketCap"`
		Binance struct {
			URL string `env-required:"true" yaml:"url"`
		} `yaml:"binance"`
		CoinAPI struct {
			URL        string `env-required:"true" yaml:"url"`
			HeaderName string `env-required:"true" yaml:"headerName"`
			APIKey     string `env-required:"true" env:"X-CoinAPI-Key"`
		} `yaml:"coinAPI"`
		Coinbase struct {
			URL string `env-required:"true" yaml:"url"`
		} `yaml:"coinbase"`
	} `yaml:"cryptoProviders"`
	EmailSending struct {
		SenderAddress string `env-required:"true" yaml:"senderAddress"`
		PublicKey     string `env-required:"true" env:"MAILJET_PUBLIC_KEY"`
		PrivateKey    string `env-required:"true" env:"MAILJET_PRIVATE_KEY"`
	} `yaml:"emailSending"`
	Database struct {
		FilePath string `env-required:"true" yaml:"filePath"`
	} `yaml:"database"`
	Cache struct {
		RateCacheTTL time.Duration `env-required:"true" yaml:"rateCacheTTL"`
	} `yaml:"cache"`
}

func GetConfig() *Config {
	var cfg = &Config{}
	var once sync.Once

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	once.Do(func() {
		err := cleanenv.ReadConfig(filepath.Join(basepath, "config.yml"), cfg)
		if err != nil {
			return
		}
	})

	return cfg
}
