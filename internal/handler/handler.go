package handler

import (
	"localx/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services     *services.Services
	tokenStorage *InMemoryTokenStorage
}

func NewHandler(s *services.Services) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// base middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	auth := router.Group("/auth")
	{
		verification := auth.Group("/verification")
		{
			//Перед отправ
			verification.POST("/sendCode", h.SendVerificationCode)
			traveler := verification.Group("/traveler")
			{
				traveler.POST("/sign-in", h.TravelerSignIn)
				//ошибка при получении всех пользователей, но это ручка не здесь будет
				traveler.GET("/", h.GetAllTraveler)
				traveler.POST("/sign-up", h.TravelerSignUp)
			}
		}

		// TODO для компании регистрацию и логин
	}

	tour := router.Group("/tour", h.userIdentity) //функция для идентификации пользователя
	{
		tour.POST("/", h.CreateTour)
		tour.GET("/:tour_id", h.GetTourById)

	}

	return router
}
