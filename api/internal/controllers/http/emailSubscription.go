package http

import (
	"api/internal/controllers/http/dto"
	"api/internal/controllers/http/response"
	"api/internal/customerrors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) sendEmails(c *gin.Context) {
	err := h.useCases.SendEmailsUseCase.SendToAll()
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.StatusResponse{Status: "sent"})
}

func (h *Handler) subscribe(c *gin.Context) {
	var input dto.SubscribeEmail

	err := c.ShouldBind(&input)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "Email field is required")
		return
	}

	err = h.useCases.SubscribeEmailUseCase.Subscribe(input.Email)
	if err != nil {
		if errors.Is(err, customerrors.ErrEmailDuplicate) {
			response.NewErrorResponse(c, http.StatusConflict, err.Error())
			return
		}

		response.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response.StatusResponse{Status: "subscribed"})
}
