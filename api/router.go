package api

import (
	"tender_bid_system/api/handler"
	"tender_bid_system/api/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	authHandler *handler.UserHandler,
	tenderHandler *handler.TenderHandler,
	bidHandler *handler.BidHandler,
	notiHandler *handler.NotificationHandler,
) *gin.Engine {
	router := gin.Default()

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	router.POST("/tender", middleware.RoleMiddleware("contractor"), tenderHandler.CreateTender)
	router.GET("/tenders", middleware.RoleMiddleware("contractor"), tenderHandler.ListTenders)
	router.PUT("/tender/:id", middleware.RoleMiddleware("contractor"), tenderHandler.UpdateTender)
	router.DELETE("/tender/:id", middleware.RoleMiddleware("contractor"), tenderHandler.DeleteTender)

	router.POST("/tenders/bids", middleware.RoleMiddleware("client"), bidHandler.SubmitBid)
	router.GET("/tenders/bids/:id", middleware.RoleMiddleware("client"), bidHandler.ViewBidsByTenderID)
	router.GET("/tenders/bids/contractor/:id", bidHandler.ViewBidsByContractorID)
	router.GET("/bid", bidHandler.GetBidsByPrice)

	router.POST("/notifications", notiHandler.CreateNotification)

	return router
}
