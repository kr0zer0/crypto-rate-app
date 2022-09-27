package http

import (
	"api/internal/customerrors"
	"api/internal/usecases"
	mock_usecases_contracts "api/internal/usecases/usecases_contracts/mocks"
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHTTPHandler_sendMails(t *testing.T) {
	type mockBehavior func(s *mock_usecases_contracts.MockSendEmailsUseCase)

	type test struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}

	testTable := []test{
		{
			name: "OK",
			mockBehavior: func(s *mock_usecases_contracts.MockSendEmailsUseCase) {
				s.EXPECT().SendToAll().Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":"sent"}`,
		},
		{
			name: "Error",
			mockBehavior: func(s *mock_usecases_contracts.MockSendEmailsUseCase) {
				s.EXPECT().SendToAll().Return(errors.New("some error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockController := gomock.NewController(t)

			emailSubMock := mock_usecases_contracts.NewMockSendEmailsUseCase(mockController)
			testCase.mockBehavior(emailSubMock)

			services := &usecases.UseCases{SendEmailsUseCase: emailSubMock}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/send", handler.sendEmails)

			responseRecorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/send", bytes.NewBufferString(""))

			r.ServeHTTP(responseRecorder, request)

			assert.Equal(t, testCase.expectedStatusCode, responseRecorder.Code)
			assert.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}

func TestHTTPHandler_subscribe(t *testing.T) {
	type mockBehavior func(s *mock_usecases_contracts.MockSubscribeEmailUseCase, email string)

	type test struct {
		name                 string
		emailInput           string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}

	testTable := []test{
		{
			name:       "OK",
			emailInput: "some.email@mail.com",
			mockBehavior: func(s *mock_usecases_contracts.MockSubscribeEmailUseCase, email string) {
				s.EXPECT().Subscribe(email).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"status":"subscribed"}`,
		},
		{
			name: "No email input",
			mockBehavior: func(s *mock_usecases_contracts.MockSubscribeEmailUseCase, email string) {
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"Email field is required"}`,
		},
		{
			name:       "Email duplicate",
			emailInput: "some.email@mail.com",
			mockBehavior: func(s *mock_usecases_contracts.MockSubscribeEmailUseCase, email string) {
				s.EXPECT().Subscribe(email).Return(customerrors.ErrEmailDuplicate)
			},
			expectedStatusCode:   http.StatusConflict,
			expectedResponseBody: fmt.Sprintf(`{"message":"%v"}`, customerrors.ErrEmailDuplicate.Error()),
		},
		{
			name:       "Some internal error",
			emailInput: "some.email@mail.com",
			mockBehavior: func(s *mock_usecases_contracts.MockSubscribeEmailUseCase, email string) {
				s.EXPECT().Subscribe(email).Return(errors.New("some error"))
			},
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"some error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockController := gomock.NewController(t)

			emailSubMock := mock_usecases_contracts.NewMockSubscribeEmailUseCase(mockController)
			testCase.mockBehavior(emailSubMock, testCase.emailInput)

			services := &usecases.UseCases{SubscribeEmailUseCase: emailSubMock}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/subscribe", handler.subscribe)

			responseRecorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/subscribe?email="+testCase.emailInput, nil)

			r.ServeHTTP(responseRecorder, request)

			assert.Equal(t, testCase.expectedStatusCode, responseRecorder.Code)
			assert.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}
