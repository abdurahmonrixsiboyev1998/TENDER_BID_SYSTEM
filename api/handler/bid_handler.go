package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"tender_bid_system/model"
	"tender_bid_system/service"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Price         float64 `json:"price"`
	Delivery_time int     `json:"delivery_time"`
}

type BidHandler struct {
	service     *service.BidService
	serviceUser *service.UserService
}

func NewBidHandler(service *service.BidService, serviceUser *service.UserService) *BidHandler {
	return &BidHandler{
		service:     service,
		serviceUser: serviceUser,
	}
}

func (h *BidHandler) SubmitBid(c *gin.Context) {
	var bid model.Bid
	if err := c.ShouldBindJSON(&bid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	role, err := h.serviceUser.GetUserByID(c.Request.Context(), bid.ContraktorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user role"})
		return
	}
	if role != "contractor" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Only contractors can submit bids"})
		return
	}

	createdBid, err := h.service.SubmitBid(c.Request.Context(), &bid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit bid"})
		return
	}
	c.JSON(http.StatusCreated, createdBid)
}

func (h *BidHandler) ViewBidsByTenderID(c *gin.Context) {
	var id string
	id = c.Param("id")
	tenderId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tender ID"})
		return
	}
	bids, err := h.service.ViewBidsByTenderID(c.Request.Context(), tenderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bids)
}

func (h *BidHandler) ViewBidsByContractorID(c *gin.Context) {
	var id string
	id = c.Param("id")
	contractorId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid contractor ID"})
		return
	}
	bids, err := h.service.ViewBidsByContractorID(c.Request.Context(), contractorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bids)
}

func (h *BidHandler) GetBidsByPrice(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bids, err := h.service.GetBidsByPrice(c.Request.Context(), req.Price, req.Delivery_time)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bids)
}
