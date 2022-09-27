package fileStorage

import (
	"api/internal/customerrors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailSubscriptionRepository_addToSorted(t *testing.T) {
	type test struct {
		name          string
		emailsList    []string
		emailToInsert string
		expectedList  []string
		expectedErr   error
	}

	testTable := []test{
		{
			name:          "Inserting to the end",
			emailsList:    []string{"a", "b", "c"},
			emailToInsert: "d",
			expectedList:  []string{"a", "b", "c", "d"},
			expectedErr:   nil,
		},
		{
			name:          "Inserting inside",
			emailsList:    []string{"a", "b", "d"},
			emailToInsert: "c",
			expectedList:  []string{"a", "b", "c", "d"},
			expectedErr:   nil,
		},
		{
			name:          "Inserting duplicate",
			emailsList:    []string{"a", "b", "c"},
			emailToInsert: "c",
			expectedList:  nil,
			expectedErr:   customerrors.ErrEmailDuplicate,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			repo := &EmailSubscriptionRepository{}
			gotList, err := repo.addToSorted(testCase.emailsList, testCase.emailToInsert)
			assert.Equal(t, testCase.expectedList, gotList)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}
