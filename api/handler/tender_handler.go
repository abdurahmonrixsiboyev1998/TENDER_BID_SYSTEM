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
