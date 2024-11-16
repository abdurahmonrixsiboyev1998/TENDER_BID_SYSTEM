package handler

import (
	"net/http"
	"tender_bid_system/model"
	"tender_bid_system/service"
	"time"

	"github.com/gin-gonic/gin"
)

type BidHandler struct {
	service *service.BidService
}

func NewBidHandler(service *service.BidService) *BidHandler {
	return &BidHandler{service: service}
}

func (h *BidHandler) SubmitBid(c *gin.Context) {
	var bid model.Bid
	if err := c.ShouldBindJSON(&bid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdBid, err := h.service.SubmitBid(c.Request.Context(), &bid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdBid)
}

func (h *BidHandler) ViewBidsByTenderID(c *gin.Context) {
	var tenderID int
	if err := c.ShouldBindJSON(&tenderID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bids, err := h.service.ViewBidsByTenderID(c.Request.Context(), tenderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bids)
}

func (h *BidHandler) GetBidsByPrice(c *gin.Context) {
	var price float64
	var delivery_time time.Time
	if err := c.ShouldBindJSON(&price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&delivery_time); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bids, err := h.service.GetBidsByPrice(c.Request.Context(), price, delivery_time)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bids)
}
