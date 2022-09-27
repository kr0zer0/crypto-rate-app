package http

import (
	"api/internal/controllers/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getCurrentExchangeRate(c *gin.Context) {
	rate, err := h.useCases.GetRateUseCase.GetBtcUahRate()
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, rate.Rate)
}
