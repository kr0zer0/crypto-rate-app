package fileStorage

import (
	"api/internal/constants"
	customerrors "api/internal/customerrors"
	"encoding/json"
	"os"
	"sort"

	"api/data"
)

type EmailSubscriptionRepository struct {
	filepath string
}

func NewEmailSubscriptionRepository(filepath string) *EmailSubscriptionRepository {
	return &EmailSubscriptionRepository{
		filepath: filepath,
	}
}

func (r *EmailSubscriptionRepository) Add(email string) error {
	emails, err := r.GetAll()
	if err != nil {
		return err
	}

	emails, err = r.addToSorted(emails, email)
	if err != nil {
		return err
	}

	records := data.SubscribedEmails{
		Emails: emails,
	}

	updatedData, err := json.Marshal(records)
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filepath, updatedData, constants.WriteFilePerm)
	if err != nil {
		return err
	}

	return nil
}

func (r *EmailSubscriptionRepository) GetAll() ([]string, error) {
	records := data.SubscribedEmails{}
	file, err := os.ReadFile(r.filepath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &records)
	if err != nil {
		return nil, err
	}
	return records.Emails, nil
}

func (r *EmailSubscriptionRepository) addToSorted(sourceSlice []string, itemToAdd string) ([]string, error) {
	index := sort.SearchStrings(sourceSlice, itemToAdd)
	if index != len(sourceSlice) {
		if sourceSlice[index] == itemToAdd {
			return nil, customerrors.ErrEmailDuplicate
		}
	}

	sourceSlice = append(sourceSlice, "")
	copy(sourceSlice[index+1:], sourceSlice[index:])
	sourceSlice[index] = itemToAdd

	return sourceSlice, nil
}
