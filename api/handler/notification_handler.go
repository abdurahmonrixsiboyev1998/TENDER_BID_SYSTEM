package handler

import (
	"net/http"
	"tender_bid_system/model"
	"tender_bid_system/service"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	service *service.NotificationService
}

func NewNotificationHandler(service *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: service}
}

// CreateNotification godoc
// @Summary Create a new notification
// @Description Create a new notification to be sent to the user
// @Tags Notification
// @Accept json
// @Produce json
// @Param notification body model.Notification true "Notification object"
// @Success 201 {object} model.Notification
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /notifications [post]
func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var notification model.Notification
	if err := c.ShouldBindJSON(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	createdNotification, err := h.service.CreateNotification(c.Request.Context(), &notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}
	c.JSON(http.StatusCreated, createdNotification)
}
