package usecases

import (
	"api/internal/usecases/usecases_contracts"
	"fmt"
)

type SendEmailsUseCase struct {
	emailSubsRepo  usecases_contracts.EmailSubscriptionRepo
	mailer         usecases_contracts.Mailer
	getRateUseCase usecases_contracts.GetRateUseCase
}

func NewSendEmailsUseCase(emailSubsRepo usecases_contracts.EmailSubscriptionRepo, mailer usecases_contracts.Mailer, cryptoService usecases_contracts.GetRateUseCase) *SendEmailsUseCase {
	return &SendEmailsUseCase{
		emailSubsRepo:  emailSubsRepo,
		mailer:         mailer,
		getRateUseCase: cryptoService,
	}
}

func (u *SendEmailsUseCase) SendToAll() error {
	emails, err := u.emailSubsRepo.GetAll()
	if err != nil {
		return err
	}

	rate, err := u.getRateUseCase.GetBtcUahRate()
	if err != nil {
		return err
	}

	err = u.mailer.SendToList(emails, fmt.Sprintf("%.2fUAH", rate))
	if err != nil {
		return err
	}

	return nil
}
