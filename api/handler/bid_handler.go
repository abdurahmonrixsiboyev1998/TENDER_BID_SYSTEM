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

// SubmitBid godoc
// @Summary Submit a new bid
// @Description Allows contractors to submit a bid for a tender
// @Tags Bid
// @Accept json
// @Produce json
// @Param bid body model.Bid true "Bid details"
// @Success 201 {object} model.Bid
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Failure 500 {object} string
// @Router /bids [post]
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

// ViewBidsByTenderID godoc
// @Summary View bids by tender ID
// @Description Retrieve all bids associated with a specific tender ID
// @Tags Bid
// @Accept json
// @Produce json
// @Param id path string true "Tender ID"
// @Success 200 {array} model.Bid "List of bids"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /tender/bids [get]
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

// ViewBidsByContractorID godoc
// @Summary View bids by contractor ID
// @Description Retrieve all bids submitted by a specific contractor using their ID
// @Tags Bid
// @Accept json
// @Produce json
// @Param id path string true "Contractor ID"
// @Success 200 {array} model.Bid "List of bids submitted by the contractor"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /trendes/bids/contractor/ [get]
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

// GetBidsByPrice godoc
// @Summary Get bids by price and delivery time
// @Description Retrieve bids filtered by price and delivery time
// @Tags Bid
// @Accept json
// @Produce json
// @Param request body Request true "Request containing price and delivery time filter"
// @Success 200 {array} model.Bid
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /bids/price [post]
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
