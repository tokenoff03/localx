package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TravelerInput struct {
	Email string `json:"email" binding required`
}

type VerifyCodeInput struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

func (h *Handler) SendVerificationCode(c *gin.Context) {
	var input TravelerInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input, need email")
		return
	}

	// Генерация кода
	code, err := h.services.AuthTraveler.GenerateAndSaveOTP(input.Email)
	if err != nil {
		fmt.Println("Error generating OTP:", err)
		return
	}

	// Отправка кода на телефон
	err = h.services.SendEmail(input.Email, "Ваш код:", code) // Нужно реализовать интеграцию с SMS API
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "verification code sent",
	})
}

func (h *Handler) TravelerSignIn(c *gin.Context) {
	var input VerifyCodeInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input, need phoneNumber")
		return
	}
	_, err := h.services.ValidateOTP(input.Email, input.Code)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.services.GetTraveler(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Email == "" {
		c.JSON(http.StatusOK, res)
		return
	}

	accessToken, err := h.services.GenerateToken(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := h.services.GenerateRefreshToken(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Сохраняем токены в In-Memory хранилище
	h.services.AuthTraveler.StoreTokens(input.Email, accessToken, refreshToken)

	// Возвращаем токены клиенту
	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
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
