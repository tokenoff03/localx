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
			verification.POST("/sendCode", h.SendVerificationCode)
			traveler := verification.Group("/traveler")
			{
				traveler.POST("/sign-in", h.TravelerSignIn)
				traveler.GET("/", h.GetAllTraveler)
			}
		}

		// company := auth.Group("/company")
		// {
		// 	company.POST("sign-in")
		// }
	}

	tour := router.Group("/tour")
	{
		tour.POST("/", h.CreateTour)
		tour.GET("/:tour_id", h.GetTourById)

	}

	return router
}
