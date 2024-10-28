package handler

import (
	"localx/internal/models"
	"localx/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTour(c *gin.Context) {
	var input models.Tour

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()

	id, err := h.services.Tour.CreateTour(ctx, input, int64(input.CompanyID))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "status": "tour created"})
}

func (h *Handler) GetTourById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id param is not correct")
		return
	}

	ctx := c.Request.Context()
	tour, err := h.services.Tour.GetTourById(ctx, int64(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tour)
}

func (h *Handler) UpdateTour(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid tour ID")
		return
	}

	var input models.Tour
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.ID = id

	ctx := c.Request.Context()
	err = h.services.Tour.UpdateTour(ctx, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "tour updated"})
}

func (h *Handler) DeleteTour(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid tour ID")
		return
	}

	ctx := c.Request.Context()
	err = h.services.Tour.DeleteTour(ctx, int64(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "tour deleted"})
}

func (h *Handler) UpdateTourDetails(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid tour ID")
		return
	}

	var input models.Tour
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	existingTour, err := h.services.Tour.GetTourById(c.Request.Context(), int64(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	updatedTour := utils.ApplyTourUpdatesDetails(input, existingTour)
	updatedTour.ID = int(id)

	ctx := c.Request.Context()
	if err := h.services.Tour.UpdateTour(ctx, updatedTour); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Tour details updated"})
}
