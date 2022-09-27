package integration

import (
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func (s *IntegrationTestSuite) TestGetCurrentRate() {
	router := s.handler.InitRouter()

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/api/rate", nil)

	router.ServeHTTP(responseRecorder, request)

	assert.Equal(s.T(), http.StatusOK, responseRecorder.Code)
}
