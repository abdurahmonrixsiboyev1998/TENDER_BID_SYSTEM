package handler

import (
	"net/http"
	"strconv"
	"tender_bid_system/model"
	"tender_bid_system/service"

	"github.com/gin-gonic/gin"
)

type TenderHandler struct {
	service *service.TenderService
}

func NewTenderHandler(service *service.TenderService) *TenderHandler {
	return &TenderHandler{service: service}
}

// CreateTender godoc
// @Summary Create a new tender
// @Description Create a new tender and store it in the system
// @Tags Tender
// @Accept json
// @Produce json
// @Param request body model.Tender true "Tender details"
// @Success 201 {object} model.Tender
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /tenders [post]
func (h *TenderHandler) CreateTender(c *gin.Context) {
	var tender model.Tender
	if err := c.ShouldBindJSON(&tender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTender, err := h.service.CreateTender(c.Request.Context(), &tender)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTender)
}

func (h *TenderHandler) ListTenders(c *gin.Context) {
	tenders, err := h.service.ListTenders(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tenders)
}

// UpdateTender godoc
// @Summary Update an existing tender
// @Description Update the details of an existing tender
// @Tags Tender
// @Accept json
// @Produce json
// @Param request body model.Tender true "Updated tender details"
// @Success 200 {object} model.Tender
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /tenders [put]
func (h *TenderHandler) UpdateTender(c *gin.Context) {
	var tender model.Tender
	if err := c.ShouldBindJSON(&tender); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTender, err := h.service.UpdateTender(c.Request.Context(), &tender)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTender)
}

// DeleteTender godoc
// @Summary Delete an existing tender
// @Description Delete a tender by its ID
// @Tags Tender
// @Accept json
// @Produce json
// @Param id path int true "Tender ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /tenders/{id} [delete]
func (h *TenderHandler) DeleteTender(c *gin.Context) {
	var id string
	id = c.Param("id")
	tenderId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tender ID"})
		return
	}
	err = h.service.DeleteTender(c.Request.Context(), tenderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tender"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tender deleted successfully"})
}
