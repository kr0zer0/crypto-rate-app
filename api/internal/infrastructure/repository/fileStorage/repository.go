package fileStorage

import (
	"api/internal/usecases/usecases_contracts"
)

func NewRepository(emailSub usecases_contracts.EmailSubscriptionRepo) *usecases_contracts.Repository {
	return &usecases_contracts.Repository{
		EmailSubscriptionRepo: emailSub,
	}
}
