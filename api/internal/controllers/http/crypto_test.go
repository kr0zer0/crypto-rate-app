package http

import (
	"api/internal/entities"
	"api/internal/usecases"
	mock_usecases_contracts "api/internal/usecases/usecases_contracts/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHTTPHandler_getCurrentExchangeRate(t *testing.T) {
	type mockBehavior func(s *mock_usecases_contracts.MockGetRateUseCase)

	type test struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}

	testTable := []test{
		{
			name: "OK",
			mockBehavior: func(s *mock_usecases_contracts.MockGetRateUseCase) {
				s.EXPECT().GetBtcUahRate().Return(entities.NewRate(entities.NewCurrencyPair(entities.BTC, entities.UAH), 777.777), nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: "777.777",
		},
		{
			name: "Error",
			mockBehavior: func(s *mock_usecases_contracts.MockGetRateUseCase) {
				s.EXPECT().GetBtcUahRate().Return(nil, errors.New("some error"))
			},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"some error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockController := gomock.NewController(t)

			cryptoServiceMock := mock_usecases_contracts.NewMockGetRateUseCase(mockController)
			testCase.mockBehavior(cryptoServiceMock)

			services := &usecases.UseCases{GetRateUseCase: cryptoServiceMock}
			handler := NewHandler(services)

			r := gin.New()
			r.GET("/rate", handler.getCurrentExchangeRate)

			responseRecorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/rate", nil)

			r.ServeHTTP(responseRecorder, request)

			assert.Equal(t, testCase.expectedStatusCode, responseRecorder.Code)
			assert.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}
