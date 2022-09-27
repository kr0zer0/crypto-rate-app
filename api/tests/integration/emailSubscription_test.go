package integration

import (
	"api/internal/customerrors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func (s *IntegrationTestSuite) TestSubscribe() {
	router := s.handler.InitRouter()

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/api/subscribe?email=email@mail.com", nil)

	router.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.Equal(s.T(), `{"status":"subscribed"}`, responseRecorder.Body.String())
}

func (s *IntegrationTestSuite) TestSubscribe_EmptyInput() {
	router := s.handler.InitRouter()

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/api/subscribe", nil)

	router.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusBadRequest, responseRecorder.Code)
	assert.Equal(s.T(), `{"message":"Email field is required"}`, responseRecorder.Body.String())
}

func (s *IntegrationTestSuite) TestSubscribe_Duplicate() {
	router := s.handler.InitRouter()

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/api/subscribe?email=email@mail.com", nil)

	router.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusConflict, responseRecorder.Code)
	assert.Equal(s.T(),
		fmt.Sprintf(`{"message":"%v"}`, customerrors.ErrEmailDuplicate.Error()),
		responseRecorder.Body.String())
}

func (s *IntegrationTestSuite) TestSendEmails() {
	s.mailerMock.EXPECT().SendToList([]string{}, gomock.Any()).Return(nil)

	router := s.handler.InitRouter()

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/api/sendEmails", nil)

	router.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
	assert.Equal(s.T(), `{"status":"sent"}`, responseRecorder.Body.String())
}
