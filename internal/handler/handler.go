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

	// auth := router.Group("/auth")
	// {
	// 	auth.POST("sign-up", h.SignUp)
	// 	auth.POST("sign-in", h.SignIn)
	// }

	tour := router.Group("/tour")
	{
		tour.POST("/", h.CreateTour)
		tour.GET("/:tour_id", h.GetTourById)

	}

	return router
}
