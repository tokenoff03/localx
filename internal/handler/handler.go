package handler

import (
	"localx/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Services
}

func NewHandler(s *services.Services) *Handler {
	return &Handler{services: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		traveler := auth.Group("/traveler")
		{
			traveler.POST("/sign-in", h.TravelerSignIn)
			traveler.GET("/", h.GetAllTraveler)
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
