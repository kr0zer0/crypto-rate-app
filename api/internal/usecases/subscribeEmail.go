package usecases

import "api/internal/usecases/usecases_contracts"

type SubscribeEmailUseCase struct {
	emailSubsRepo usecases_contracts.EmailSubscriptionRepo
}

func NewSubscribeEmailUseCase(emailSubsRepo usecases_contracts.EmailSubscriptionRepo) *SubscribeEmailUseCase {
	return &SubscribeEmailUseCase{emailSubsRepo: emailSubsRepo}
}

func (u *SubscribeEmailUseCase) Subscribe(email string) error {
	err := u.emailSubsRepo.Add(email)
	if err != nil {
		return err
	}

	return nil
}
