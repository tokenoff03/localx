package handler

import (
	"encoding/json"
	"errors"
	"localx/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTour(c *gin.Context) {
	time.Sleep(5 * time.Second)

	var input models.Tour

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx := c.Request.Context()

	id, err := h.services.Tour.CreateTour(ctx, input, input.CompanyID)
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
	tour, err := h.services.Tour.GetTourById(ctx, id)
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
	err = h.services.Tour.DeleteTour(ctx, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "tour deleted"})
}

func (h *Handler) UpdateTourField(c *gin.Context) {
	idStr := c.Param("id")
	field := c.Param("field")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid tour ID")
		return
	}

	fieldUpdaters := map[string]func(ctx *gin.Context, id int) error{
		"title":                  h.updateTitle,
		"start_time":             h.updateStartTime,
		"end_time":               h.updateEndTime,
		"languages":              h.updateLanguages,
		"free_cancellation":      h.updateFreeCancellation,
		"cancellation_condition": h.updateCancellationCondition,
		"description":            h.updateDescription,
		"meeting_place":          h.updateMeetingPlace,
		"arrival_place":          h.updateArrivalPlace,
		"what_is_included":       h.updateWhatIsIncluded,
		"what_to_prepare":        h.updateWhatToPrepare,
		"prohibitions":           h.updateProhibitions,
		"price":                  h.updatePrice,
		"images":                 h.updateImages,
	}

	updater, exists := fieldUpdaters[field]
	if !exists {
		newErrorResponse(c, http.StatusBadRequest, "Invalid field name")
		return
	}

	err = updater(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": field + " updated"})
}

func (h *Handler) updateTitle(c *gin.Context, id int) error {
	var input struct {
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetTitle(ctx, id, input.Title)
}

func (h *Handler) updateStartTime(c *gin.Context, id int) error {
	var input struct {
		StartTime string `json:"start_time" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	startTime, err := time.Parse(time.RFC3339, input.StartTime)
	if err != nil {
		return errors.New("invalid start_time format")
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetStartTime(ctx, id, startTime)
}

func (h *Handler) updateEndTime(c *gin.Context, id int) error {
	var input struct {
		EndTime string `json:"end_time" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	endTime, err := time.Parse(time.RFC3339, input.EndTime)
	if err != nil {
		return errors.New("invalid end_time format")
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetEndTime(ctx, id, endTime)
}

func (h *Handler) updateLanguages(c *gin.Context, id int) error {
	var input struct {
		Languages string `json:"languages" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetLanguages(ctx, id, input.Languages)
}

func (h *Handler) updateFreeCancellation(c *gin.Context, id int) error {
	var input struct {
		FreeCancellation bool `json:"free_cancellation"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetFreeCancellation(ctx, id, input.FreeCancellation)
}

func (h *Handler) updateCancellationCondition(c *gin.Context, id int) error {
	var input struct {
		CancellationCondition map[string]interface{} `json:"cancellation_condition" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	cancellationConditionJSON, err := json.Marshal(input.CancellationCondition)
	if err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetCancellationCondition(ctx, id, string(cancellationConditionJSON))
}

func (h *Handler) updateDescription(c *gin.Context, id int) error {
	var input struct {
		Description string `json:"description" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetDescription(ctx, id, input.Description)
}

func (h *Handler) updateMeetingPlace(c *gin.Context, id int) error {
	var input struct {
		MeetingPlace string `json:"meeting_place" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetMeetingPlace(ctx, id, input.MeetingPlace)
}

func (h *Handler) updateArrivalPlace(c *gin.Context, id int) error {
	var input struct {
		ArrivalPlace string `json:"arrival_place" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetArrivalPlace(ctx, id, input.ArrivalPlace)
}

func (h *Handler) updateWhatIsIncluded(c *gin.Context, id int) error {
	var input struct {
		WhatIsIncluded string `json:"what_is_included" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetWhatIsIncluded(ctx, id, input.WhatIsIncluded)
}

func (h *Handler) updateWhatToPrepare(c *gin.Context, id int) error {
	var input struct {
		WhatToPrepare string `json:"what_to_prepare" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetWhatToPrepare(ctx, id, input.WhatToPrepare)
}

func (h *Handler) updateProhibitions(c *gin.Context, id int) error {
	var input struct {
		Prohibitions string `json:"prohibitions" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetProhibitions(ctx, id, input.Prohibitions)
}

func (h *Handler) updatePrice(c *gin.Context, id int) error {
	var input struct {
		Price int `json:"price" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetPrice(ctx, id, input.Price)
}

func (h *Handler) updateImages(c *gin.Context, id int) error {
	var input struct {
		Images string `json:"images" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		return err
	}

	ctx := c.Request.Context()
	return h.services.Tour.SetImages(ctx, id, input.Images)
}
