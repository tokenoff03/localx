package handler

import (
	"fmt"
	"net/http"

	"localx/internal/models"

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

	// Отправка кода на почту
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
	//Проверка кода
	_, err := h.services.ValidateOTP(input.Email, input.Code)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Запршавыем почту с базы
	res, err := h.services.GetTravelerByEmail(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	//В сервисе есть проверка на ошибку, если нет в базе то возвращем пустой обьект. Фронт проверяет на пустоты, если пусто то кидает запрос на регистрацию
	if res.Email == "" {
		c.JSON(http.StatusOK, res)
		return
	}
	//Если есть почта то создаем токены
	accessToken, err := h.services.GenerateToken(res.ID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := h.services.GenerateRefreshToken(res.ID)
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

func (h *Handler) TravelerSignUp(c *gin.Context) {
	var input models.TravelerSignUp
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(input)
		newErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}

	id, err := h.services.CreateTraveler(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllTraveler(c *gin.Context) {
	res, err := h.services.AuthTraveler.GetAllTraveler()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"travelers": res,
	})
}
