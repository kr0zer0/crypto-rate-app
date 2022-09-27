package usecases_contracts

import "api/internal/entities"

//go:generate mockgen -source=contracts.go -destination=mocks/mocks.go

type (
	GetRateUseCase interface {
		GetBtcUahRate() (*entities.Rate, error)
	}

	SendEmailsUseCase interface {
		SendToAll() error
	}

	SubscribeEmailUseCase interface {
		Subscribe(email string) error
	}
)

type (
	EmailSubscriptionRepo interface {
		Add(email string) error
		GetAll() ([]string, error)
	}

	Repository struct {
		EmailSubscriptionRepo
	}
)

type (
	Mailer interface {
		SendToList(emails []string, message string) error
	}

	CryptoProvider interface {
		GetExchangeRate(currencyPair entities.CurrencyPair) (*entities.Rate, error)
	}
)
