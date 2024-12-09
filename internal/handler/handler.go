package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "localx/cmd/docs"
	"localx/internal/services"
	"net/http"
	"strconv"
)

type Handler struct {
	services      *services.Services
	avatarService *services.AvatarService
}

// NewHandler создает новый экземпляр Handler с зависимыми сервисами
func NewHandler(s *services.Services, avatarService *services.AvatarService) *Handler {
	return &Handler{services: s, avatarService: avatarService}

}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
				traveler.POST("/:userID/avatar", h.UploadAvatar)
				traveler.GET("/:userID/avatar", h.GetAvatar)
				traveler.DELETE("/:userID/avatar", h.DeleteAvatar)
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

// @Summary Upload user avatar
// @Description Upload a new avatar for a user by userID.
// @Tags Avatar
// @Accept multipart/form-data
// @Produce json
// @Param userID path int true "User ID"
// @Param avatar formData file true "Avatar file"
// @Success 200 {object} map[string]string "url"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /auth/verification/traveler/{userID}/avatar [post]
func (h *Handler) UploadAvatar(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	_, err = h.avatarService.GetAvatarURL(userID)
	if err == nil {
		err = h.avatarService.DeleteAvatar(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete existing avatar"})
			return
		}
	}

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

	url, err := h.avatarService.UploadAvatar(userID, src, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}

// GetAvatar godoc
// @Summary Get user avatar
// @Description Get the avatar URL for a specific user by userID.
// @Tags Avatar
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} map[string]string "url"
// @Failure 400 {object} map[string]string "error"
// @Failure 404 {object} map[string]string "error"
// @Router /auth/verification/traveler/{userID}/avatar [get]
func (h *Handler) GetAvatar(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	url, err := h.avatarService.GetAvatarURL(userID)
	if err != nil {
		if err == services.ErrAvatarNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "avatar not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve avatar"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

// DeleteAvatar godoc
// @Summary Delete user avatar
// @Description Delete the avatar for a specific user by userID.
// @Tags Avatar
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {object} map[string]string "message"
// @Failure 400 {object} map[string]string "error"
// @Failure 404 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /auth/verification/traveler/{userID}/avatar [delete]
func (h *Handler) DeleteAvatar(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err = h.avatarService.DeleteAvatar(userID)
	if err != nil {
		if err == services.ErrAvatarNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "avatar not found"})
		} else {
			fmt.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete avatar"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "avatar deleted successfully"})
}
