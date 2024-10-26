package handler

import (
	"github.com/gin-gonic/gin"
	"localx/internal/services"
	"net/http"
)

type Handler struct {
	services      *services.Services
	avatarService *services.AvatarService
}

func NewHandler(s *services.Services, avatarService *services.AvatarService) *Handler {
	return &Handler{services: s, avatarService: avatarService}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
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
				traveler.POST("/sign-up", h.TravelerSignUp)
				traveler.POST("/upload-avatar", h.UploadAvatar)
			}
		}
	}

	tour := router.Group("/tour", h.userIdentity)
	{
		tour.POST("/", h.CreateTour)
		tour.GET("/:tour_id", h.GetTourById)
	}
	return router
}

// UploadAvatar обрабатывает загрузку аватара пользователя
func (h *Handler) UploadAvatar(c *gin.Context) {
	// Пример получения ID пользователя (в реальности это можно получить из токена или базы данных)
	userID := 1 // В будущем замените на актуальный способ получения ID пользователя

	// Получаем файл из запроса
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to open the file"})
		return
	}
	defer src.Close()

	// Загружаем аватар в Supabase
	url, err := h.avatarService.UploadAvatar(userID, src, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем URL загруженного аватара
	c.JSON(http.StatusOK, gin.H{"url": url})
}
