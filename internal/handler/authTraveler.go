package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TravelerInput struct {
	PhoneNumber string `json:"phoneNumber" binding required`
}

func (h *Handler) TravelerSignIn(c *gin.Context) {
	var input TravelerInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input, need phoneNumber")
		return
	}

	res, err := h.services.GetTraveler(input.PhoneNumber)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if res.PhoneNumber == "" {
		c.JSON(http.StatusOK, res)
		return
	}

	token, err := h.services.GenerateToken(input.PhoneNumber)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}

func (h *Handler) GetAllTraveler(c *gin.Context) {
	res, err := h.services.AuthTraveler.GetAllTraveler()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"travelers": res,
	})
}
