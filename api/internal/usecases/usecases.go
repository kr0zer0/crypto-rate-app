package usecases

import (
	"api/internal/usecases/usecases_contracts"
)

type UseCases struct {
	usecases_contracts.GetRateUseCase
	usecases_contracts.SendEmailsUseCase
	usecases_contracts.SubscribeEmailUseCase
}

func NewUseCases(getRate usecases_contracts.GetRateUseCase, sendEmails usecases_contracts.SendEmailsUseCase, subscribeEmails usecases_contracts.SubscribeEmailUseCase) *UseCases {
	return &UseCases{
		GetRateUseCase:        getRate,
		SendEmailsUseCase:     sendEmails,
		SubscribeEmailUseCase: subscribeEmails,
	}
}
