package handler

import (
	"localx/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTourById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("tour_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id param is not correct")
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) CreateTour(c *gin.Context) {
	var input models.Tour

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, input)
}
